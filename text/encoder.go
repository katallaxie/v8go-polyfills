package text

import (
	v8 "github.com/katallaxie/v8go"
)

// Decoder ...
type Decoder struct{}

// NewDecoder ...
func NewDecoder() *Decoder {
	return &Decoder{}
}

// GetDecodeFunctionCallback ...
func (d *Decoder) GetDecodeFunctionCallback() v8.FunctionCallback {
	return func(_ *v8.FunctionCallbackInfo) *v8.Value {
		return nil
	}
}

// Inject ...
func (d *Decoder) Inject(_ *v8.Isolate, _ *v8.ObjectTemplate) error {
	return nil
}
