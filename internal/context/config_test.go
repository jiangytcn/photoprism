package context

import (
	"testing"

	"github.com/photoprism/photoprism/internal/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	ctx := CliTestContext()

	assert.True(t, ctx.IsSet("assets-path"))
	assert.False(t, ctx.Bool("debug"))

	c := NewConfig(ctx)

	assert.IsType(t, new(Config), c)

	assert.Equal(t, fsutil.ExpandedFilename("../../assets"), c.AssetsPath)
	assert.False(t, c.Debug)
	assert.False(t, c.ReadOnly)
}

func TestConfig_SetValuesFromFile(t *testing.T) {
	c := NewConfig(CliTestContext())

	err := c.SetValuesFromFile("testdata/config.yml")

	assert.Nil(t, err)

	assert.False(t, c.Debug)
	assert.False(t, c.ReadOnly)
	assert.Equal(t, "/srv/photoprism", c.AssetsPath)
	assert.Equal(t, "/srv/photoprism/cache", c.CachePath)
	assert.Equal(t, "/srv/photoprism/photos/originals", c.OriginalsPath)
	assert.Equal(t, "/srv/photoprism/photos/import", c.ImportPath)
	assert.Equal(t, "/srv/photoprism/photos/export", c.ExportPath)
	assert.Equal(t, "internal", c.DatabaseDriver)
	assert.Equal(t, "root:photoprism@tcp(localhost:4000)/photoprism?parseTime=true", c.DatabaseDsn)
	assert.Equal(t, 81, c.HttpServerPort)
}
