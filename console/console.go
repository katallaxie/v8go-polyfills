package console

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/katallaxie/v8go-polyfills/runtime"

	v8 "github.com/katallaxie/v8go"
)

var _ runtime.Polyfill = (*Console)(nil)

// Opt ...
type Opt func(*Console)

// WithOutput ...
func WithOutput(output io.Writer) Opt {
	return func(c *Console) {
		c.out = output
	}
}

// New ...
func New(opt ...Opt) *Console {
	c := new(Console)
	c.out = os.Stdout

	for _, o := range opt {
		o(c)
	}

	return c
}

// Console ...
type Console struct {
	out io.Writer
}

// GetMethodName ...
func (c *Console) GetMethodName() string {
	return "log"
}

// AddTo ...
func AddTo(ctx *v8.Context, opt ...Opt) error {
	if ctx == nil {
		return errors.New("v8-polyfills/console: ctx is required")
	}

	c := New(opt...)

	iso := ctx.Isolate()
	con := v8.NewObjectTemplate(iso)

	logFn := v8.NewFunctionTemplate(iso, c.GetFunctionCallback())

	if err := con.Set(c.GetMethodName(), logFn, v8.ReadOnly); err != nil {
		return fmt.Errorf("v8-polyfills/console: %w", err)
	}

	conObj, err := con.NewInstance(ctx)
	if err != nil {
		return fmt.Errorf("v8-polyfills/console: %w", err)
	}

	global := ctx.Global()

	if err := global.Set("console", conObj); err != nil {
		return fmt.Errorf("v8-polyfills/console: %w", err)
	}

	return nil
}

// GetFunctionCallback ...
func (c *Console) GetFunctionCallback() v8.FunctionCallback {
	return func(info *v8.FunctionCallbackInfo) *v8.Value {
		if args := info.Args(); len(args) > 0 {
			inputs := make([]interface{}, len(args))
			for i, input := range args {
				inputs[i] = input
			}

			fmt.Fprintln(c.out, inputs...)
		}

		return nil
	}
}
