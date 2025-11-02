package console

import (
	"errors"
	"fmt"
	"io"
	"os"

	v8 "github.com/katallaxie/v8go"
)

// Opt is a functional option for configuring the Console.
type Opt func(*Console)

// WithOutput is an Opt that sets the output writer for the Console.
func WithOutput(output io.Writer) Opt {
	return func(c *Console) {
		c.out = output
	}
}

// Console is a polyfill for the console object.
type Console struct {
	out io.Writer
}

// GetMethodName returns the method name.
func (c *Console) GetMethodName() string {
	return "log"
}

// Inject ...
func Inject(ctx *v8.Context, _ *v8.Isolate, _ *v8.ObjectTemplate) error {
	if ctx == nil {
		return errors.New("v8-polyfills/console: ctx is required")
	}

	c := New(WithOutput(os.Stdout))

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

// Add ...
func Add(ctx *v8.Context, opts ...Opt) error {
	if ctx == nil {
		return errors.New("v8-polyfills/console: ctx is required")
	}

	c := New(opts...)

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

// New ...
func New(opt ...Opt) *Console {
	c := new(Console)
	c.out = os.Stdout

	for _, o := range opt {
		o(c)
	}

	return c
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
