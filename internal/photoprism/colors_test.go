package photoprism

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/photoprism/photoprism/internal/context"
	"github.com/stretchr/testify/assert"
)

func TestMediaFile_Colors_Testdata(t *testing.T) {
	/*
		TODO: Add and compare other images in "testdata/"
	*/
	expected := map[string]ColorPerception{
		"testdata/sharks_blue.jpg": {
			Colors:     IndexedColors{0x6, 0x6, 0x5, 0x4, 0x4, 0x4, 0x4, 0x4, 0x4},
			MainColor:  4,
			Luminance:  LightMap{0x9, 0x8, 0x7, 0x6, 0x6, 0x6, 0x5, 0x5, 0x5},
			Saturation: 14,
		},
		"testdata/cat_black.jpg": {
			Colors:     IndexedColors{0x1, 0x2, 0x2, 0x2, 0x1, 0x1, 0x3, 0x2, 0x2},
			MainColor:  2,
			Luminance:  LightMap{0x8, 0x8, 0x8, 0x8, 0x4, 0x6, 0xd, 0xc, 0x8},
			Saturation: 2,
		},
		"testdata/cat_brown.jpg": {
			Colors:     IndexedColors{0x1, 0x2, 0x2, 0x1, 0x2, 0x1, 0x1, 0x1, 0x1},
			MainColor:  2,
			Luminance:  LightMap{0x5, 0x9, 0x8, 0x7, 0xb, 0x7, 0x3, 0x6, 0x7},
			Saturation: 2,
		},
		"testdata/cat_yellow_grey.jpg": {
			Colors:     IndexedColors{0x1, 0x1, 0x2, 0x1, 0x2, 0x2, 0x1, 0x1, 0xa},
			MainColor:  2,
			Luminance:  LightMap{0x5, 0x6, 0x8, 0x6, 0x8, 0x8, 0x5, 0x5, 0x6},
			Saturation: 4,
		},
	}

	err := filepath.Walk("testdata", func(filename string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if fileInfo.IsDir() || strings.HasPrefix(filepath.Base(filename), ".") {
			return nil
		}

		mediaFile, err := NewMediaFile(filename)

		if err != nil || !mediaFile.IsJpeg() {
			return nil
		}

		t.Run(filename, func(t *testing.T) {
			p, err := mediaFile.Colors()

			t.Log(p, err)

			assert.Nil(t, err)
			assert.True(t, p.Saturation.Int() >= 0)
			assert.True(t, p.Saturation.Int() < 16)
			assert.NotEmpty(t, p.MainColor.Name())

			if e, ok := expected[filename]; ok {
				assert.Equal(t, e, p)
			}
		})

		return nil
	})

	if err != nil {
		t.Log(err.Error())
	}
}

func TestMediaFile_Colors(t *testing.T) {
	ctx := context.TestContext()

	ctx.InitializeTestData(t)

	t.Run("dog.jpg", func(t *testing.T) {
		if mediaFile, err := NewMediaFile(ctx.ImportPath() + "/dog.jpg"); err == nil {
			p, err := mediaFile.Colors()

			t.Log(p, err)

			assert.Nil(t, err)
			assert.Equal(t, 2, p.Saturation.Int())
			assert.IsType(t, IndexedColors{}, p.Colors)
			assert.Equal(t, "grey", p.MainColor.Name())
			assert.Equal(t, IndexedColors{0x1, 0x2, 0x1, 0x2, 0x2, 0x1, 0x1, 0x1, 0x0}, p.Colors)
			assert.Equal(t, LightMap{0x5, 0x9, 0x7, 0xa, 0x8, 0x5, 0x5, 0x6, 0x2}, p.Luminance)
		} else {
			t.Error(err)
		}
	})

	t.Run("ape.jpeg", func(t *testing.T) {
		if mediaFile, err := NewMediaFile(ctx.ImportPath() + "/ape.jpeg"); err == nil {
			p, err := mediaFile.Colors()

			t.Log(p, err)

			assert.Nil(t, err)
			assert.Equal(t, 2, p.Saturation.Int())
			assert.IsType(t, IndexedColors{}, p.Colors)
			assert.Equal(t, "teal", p.MainColor.Name())
			assert.Equal(t, IndexedColors{0x8, 0x8, 0x2, 0x8, 0x2, 0x1, 0x8, 0x1, 0x2}, p.Colors)
			assert.Equal(t, LightMap{0x7, 0x7, 0x6, 0x7, 0x7, 0x5, 0x7, 0x6, 0x8}, p.Luminance)
		} else {
			t.Error(err)
		}
	})

	t.Run("iphone/IMG_6788.JPG", func(t *testing.T) {
		if mediaFile, err := NewMediaFile(ctx.ImportPath() + "/iphone/IMG_6788.JPG"); err == nil {
			p, err := mediaFile.Colors()

			t.Log(p, err)

			assert.Nil(t, err)
			assert.Equal(t, 2, p.Saturation.Int())
			assert.IsType(t, IndexedColors{}, p.Colors)
			assert.Equal(t, "grey", p.MainColor.Name())
			assert.Equal(t, IndexedColors{0x2, 0x1, 0x2, 0x1, 0x1, 0x1, 0x2, 0x1, 0x2}, p.Colors)
		} else {
			t.Error(err)
		}
	})

	t.Run("raw/20140717_154212_1EC48F8489.jpg", func(t *testing.T) {
		if mediaFile, err := NewMediaFile(ctx.ImportPath() + "/raw/20140717_154212_1EC48F8489.jpg"); err == nil {
			p, err := mediaFile.Colors()

			t.Log(p, err)

			assert.Nil(t, err)
			assert.Equal(t, 2, p.Saturation.Int())
			assert.IsType(t, IndexedColors{}, p.Colors)
			assert.Equal(t, "grey", p.MainColor.Name())

			assert.Equal(t, IndexedColors{0x3, 0x2, 0x2, 0x1, 0x2, 0x2, 0x2, 0x2, 0x1}, p.Colors)
		} else {
			t.Error(err)
		}
	})
}
