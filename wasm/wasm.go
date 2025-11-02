package wasm

import (
	v8 "github.com/katallaxie/v8go"
)

// Option ...
type Option func(*Module)

// Module ...
type Module struct {
	ModulePath string
}

// New ...
func New(opts ...Option) *Module {
	m := new(Module)

	for _, opt := range opts {
		opt(m)
	}

	return m
}

// WithModulePath ...
func WithModulePath(path string) Option {
	return func(m *Module) {
		m.ModulePath = path
	}
}

// Inject ...
func (m *Module) Inject(_ *v8.Isolate, _ *v8.ObjectTemplate) error {
	return nil
}
