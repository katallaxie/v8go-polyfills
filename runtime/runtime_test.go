package runtime_test

import (
	"testing"

	"github.com/katallaxie/v8go-polyfills/runtime"

	v8 "github.com/katallaxie/v8go"
	"github.com/stretchr/testify/require"
)

func Test2024_10_01(t *testing.T) {
	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)

	ctx := v8.NewContext(iso, global)

	build := runtime.Compatibility["2024-10-01"]

	for _, builder := range build {
		err := builder(ctx, iso)
		require.NoError(t, err)
	}

	defer ctx.Close()

	_, err := ctx.RunScript("console.log('hello world')", "console.js")
	require.NoError(t, err)
}

func Benchmark2024_10_01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iso := v8.NewIsolate()
		defer iso.Dispose()

		global := v8.NewObjectTemplate(iso)

		ctx := v8.NewContext(iso, global)
		defer ctx.Close()

		build := runtime.Compatibility["2024-10-01"]

		for _, builder := range build {
			err := builder(ctx, iso)
			if err != nil {
				b.Fatal(err)
			}
		}

		_, err := ctx.RunScript("console.log('hello world')", "console.js")
		if err != nil {
			b.Fatal(err)
		}
	}
}
