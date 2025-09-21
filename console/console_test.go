package console_test

import (
	"os"
	"testing"

	"github.com/katallaxie/v8go-polyfills/console"

	v8 "github.com/katallaxie/v8go"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)

	ctx := v8.NewContext(iso, global)

	err := console.Add(ctx, console.WithOutput(os.Stdout))
	require.NoError(t, err)

	defer ctx.Close()

	_, err = ctx.RunScript("console.log('hello world')", "console.js")
	require.NoError(t, err)
}
