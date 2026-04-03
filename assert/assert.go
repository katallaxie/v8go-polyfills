package assert

import (
	"fmt"

	v8 "github.com/katallaxie/v8go"
)

// Assert ...
type Assert struct{}

// New ...
func New() *Assert {
	return &Assert{}
}

// Inject implements the Injector interface.
func Inject(ctx *v8.Context, iso *v8.Isolate, global *v8.ObjectTemplate) error {
	assert := New()

	fn := v8.NewFunctionTemplate(iso, assert.GetFunctionCallBack())

	if err := global.Set("assert", fn, v8.ReadOnly); err != nil {
		return fmt.Errorf("v8-polyfills/assert: %w", err)
	}

	return nil
}

// GetFunctionCallBack ...
func (b *Assert) GetFunctionCallBack() v8.FunctionCallback {
	return func(info *v8.FunctionCallbackInfo) *v8.Value {
		args := info.Args()

		if len(args) == 0 {
			return nil
		}

		ok := args[0].Boolean()
		if !ok {
			panic("assertion failed")
		}

		return nil
	}
}
