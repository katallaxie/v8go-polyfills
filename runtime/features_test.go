package runtime_test

import (
	"testing"

	v8 "github.com/katallaxie/v8go"
	"github.com/katallaxie/v8go-polyfills/assert"
	"github.com/katallaxie/v8go-polyfills/console"
	"github.com/stretchr/testify/require"

	_ "embed"
)

//go:embed features_test.js
var test string

func TestFeatures(t *testing.T) {
	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)

	err := assert.Inject(nil, iso, global)
	require.NoError(t, err)

	ctx := v8.NewContext(iso, global)

	err = console.Inject(ctx, nil, nil)
	require.NoError(t, err)

	check, err := ctx.RunScript(test, "features_test.js")
	require.NoError(t, err)
	require.False(t, check.Boolean())
}
