package root

import (
	"context"
	"testing"

	"github.com/gopasspw/gopass/internal/backend"
	"github.com/gopasspw/gopass/internal/config"
	"github.com/gopasspw/gopass/pkg/ctxutil"
	"github.com/gopasspw/gopass/tests/gptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	t.Parallel()

	u := gptest.NewUnitTester(t)
	defer u.Remove()

	ctx := context.Background()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithHidden(ctx, true)
	ctx = backend.WithCryptoBackend(ctx, backend.Plain)

	cfg := config.NewNoWrites()
	require.NoError(t, cfg.SetPath(u.StoreDir("rs")))
	rs := New(cfg)

	inited, err := rs.IsInitialized(ctx)
	assert.NoError(t, err)
	assert.False(t, inited)
	assert.NoError(t, rs.Init(ctx, "", u.StoreDir("rs"), "0xDEADBEEF"))

	inited, err = rs.IsInitialized(ctx)
	require.NoError(t, err)
	assert.True(t, inited)
	assert.NoError(t, rs.Init(ctx, "rs2", u.StoreDir("rs2"), "0xDEADBEEF"))
}
