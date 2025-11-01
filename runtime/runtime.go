package runtime

import (
	v8 "github.com/katallaxie/v8go"
	"github.com/katallaxie/v8go-polyfills/console"
)

// Polyfill is an interface that represents a polyfill.
type Polyfill interface {
	// GetFunctionCallback returns the function callback.
	GetFunctionCallback() v8.FunctionCallback
	// GetMethodName returns the method name.
	GetMethodName() string
}

// CompatibilityDate is a string that represents the date of the last update.
type CompatibilityDate string

// CompatibilityFlag is a string that represents the name of a compatibility flag.
type CompatibilityFlag string

// CompatibilityFlags is a map of compatibility flags.
type CompatibilityFlags map[CompatibilityFlag]Injector

// CompatibilityMatrix is a map of compatibility dates.
type CompatibilityMatrix map[CompatibilityDate]CompatibilityFlags

// Compatibility is a map of compatibility matrices.
var Compatibility = CompatibilityMatrix{
	"2024-10-01": {
		"console":          console.Build,
		"addEventListener": Unimplemented,
	},
}

// Injector is a function that builds a polyfill.
type Injector func(ctx *v8.Context, iso *v8.Isolate) error

// Unimplemented is a placeholder for unimplemented polyfills.
func Unimplemented(ctx *v8.Context, iso *v8.Isolate) error {
	return nil
}
