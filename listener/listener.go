package listeners

import (
	"fmt"
	"sync"

	"github.com/katallaxie/v8go-polyfills/runtime"

	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/pkg/logx"
	"github.com/katallaxie/pkg/utilx"
	v8 "github.com/katallaxie/v8go"
)

var _ runtime.Polyfill = (*Listener)(nil)

// Error ...
type Error struct {
	// Message is the error message.
	Message string
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("v8-polyfills/listener: %s", e.Message)
}

// NewDefaultError ...
func NewDefaultError() *Error {
	return &Error{Message: "an error occurred while calling the event listener"}
}

// NewError ...
func NewError(message string) *Error {
	return &Error{Message: message}
}

// NewErrorIsolateRequired ...
func NewErrorIsolateRequired() *Error {
	return &Error{Message: "isolate is required"}
}

// Opt is a functional option for configuring the listener.
type Opt func(*Listener)

// WithLogger ...
func WithLogger(l logx.Logger) Opt {
	return func(c *Listener) {
		c.log = l
	}
}

// Listener is a polyfill for the addEventListener method.
type Listener struct {
	log logx.Logger
	in  sync.Map
	out sync.Map
}

// GetMethodName ...
func (l *Listener) GetMethodName() string {
	return "addEventListener"
}

// New ...
func New(opt ...Opt) *Listener {
	c := new(Listener)
	c.log = logx.LogSink
	c.in = sync.Map{}
	c.out = sync.Map{}

	for _, o := range opt {
		o(c)
	}

	return c
}

// WithEvents ...
func WithEvents(name string, in chan *v8.Object, out chan *v8.Value) Opt {
	return func(l *Listener) {
		l.in.Store(name, in)
		l.out.Store(name, out)
	}
}

// Add ...
func Add(iso *v8.Isolate, global *v8.ObjectTemplate, opts ...Opt) error {
	if utilx.IsNil(iso) {
		return NewErrorIsolateRequired()
	}

	l := New(opts...)

	ctxFn := v8.NewFunctionTemplate(iso, l.GetFunctionCallback())

	if err := global.Set(l.GetMethodName(), ctxFn, v8.ReadOnly); err != nil {
		return NewError(err.Error())
	}

	return nil
}

// GetFunctionCallback ...
func (l *Listener) GetFunctionCallback() v8.FunctionCallback {
	return func(info *v8.FunctionCallbackInfo) *v8.Value {
		ctx := info.Context()
		args := info.Args()

		if len(args) <= 1 {
			err := NewError(fmt.Sprintf("expected 2 arguments, got %d", len(args)))

			return newErrorValue(ctx, err)
		}

		fn, err := args[1].AsFunction()
		if err != nil {
			err := NewError(fmt.Sprintf("v8-polyfills/listener: %v", err))

			return newErrorValue(ctx, err)
		}

		chn, ok := l.in.Load(conv.String(args[0]))
		if !ok {
			err := NewError(fmt.Sprintf("v8-polyfills/listener: event %s not found", args[0].String()))

			return newErrorValue(ctx, err)
		}

		rchn, ok := chn.(chan *v8.Object)
		if !ok {
			err := NewError(fmt.Sprintf("v8-polyfills/listener: event %s is not a channel", args[0].String()))

			return newErrorValue(ctx, err)
		}

		go func(chn chan *v8.Object, fn *v8.Function) {
			for e := range chn {
				v, err := fn.Call(ctx.Global(), e)
				if err != nil {
					l.log.Errorf("v8-polyfills/listener: %v", err)
				}

				out, ok := l.out.Load(conv.String(args[0]))
				if !ok {
					l.log.Errorf("v8-polyfills/listener: out channel not found")
				}

				vchn, ok := out.(chan *v8.Value)
				if !ok {
					return
				}

				vchn <- v
			}
		}(rchn, fn)

		return v8.Undefined(ctx.Isolate()) //nolint:for
	}
}

func newErrorValue(ctx *v8.Context, err error) *v8.Value {
	iso := ctx.Isolate()
	e, _ := v8.NewValue(iso, fmt.Sprintf("addListener: %v", err))

	return e
}
