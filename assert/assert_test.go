package assert

import (
	"testing"

	v8 "github.com/katallaxie/v8go"

	"github.com/stretchr/testify/assert"
)

func TestAssertOk(t *testing.T) {
	ctx, err := newV8goContext()
	assert.NoError(t, err)
	defer ctx.Close()

	_, err = ctx.RunScript("assert(Map !== \"undefined\")", "test.js")
	assert.NoError(t, err)
}

func TestAssertFail(t *testing.T) {
	ctx, err := newV8goContext()
	assert.NoError(t, err)
	defer ctx.Close()

	assert.PanicsWithValue(t, "assertion failed", func() {
		_, _ = ctx.RunScript("assert(Map === \"undefined\")", "test.js")
	})
}

func newV8goContext() (*v8.Context, error) {
	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)

	ctx := v8.NewContext(iso, global)

	if err := Inject(ctx, iso, global); err != nil {
		return nil, err
	}

	return v8.NewContext(iso, global), nil
}
