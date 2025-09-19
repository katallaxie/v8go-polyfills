package utils_test

import (
	"testing"

	"github.com/katallaxie/v8go"
	"github.com/katallaxie/v8go-polyfills/utils"

	"github.com/stretchr/testify/assert"
)

func TestNewInt32Value(t *testing.T) {
	iso := v8go.NewIsolate()
	defer iso.Dispose()

	ctx := v8go.NewContext(iso)
	defer ctx.Close()

	v, err := utils.NewInt32Value(ctx, 123)
	assert.NoError(t, err)

	assert.Equal(t, int32(123), v.Int32())
}
