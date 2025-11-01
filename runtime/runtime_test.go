package runtime_test

import (
	"testing"

	"github.com/katallaxie/v8go-polyfills/runtime"

	v8 "github.com/katallaxie/v8go"
	"github.com/stretchr/testify/require"
)

func Test_2024_10_01(t *testing.T) {
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
