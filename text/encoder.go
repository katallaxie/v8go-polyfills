package text

import (
	"github.com/katallaxie/v8go-polyfills/utils"

	v8 "github.com/katallaxie/v8go"
)

// Decoder ...
type Decoder struct {
	utils.Injector
}

// NewDecoder ...
func NewDecoder() *Decoder {
	return &Decoder{}
}

// GetDecodeFunctionCallback ...
func (d *Decoder) GetDecodeFunctionCallback() v8.FunctionCallback {
	return func(info *v8.FunctionCallbackInfo) *v8.Value {
		return nil
	}
}

// Inject ...
func (d *Decoder) Inject(iso *v8.Isolate, global *v8.ObjectTemplate) error {
	return nil
}
