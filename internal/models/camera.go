package models

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

// Camera model and make (as extracted from EXIF metadata)
type Camera struct {
	gorm.Model
	CameraSlug        string
	CameraModel       string
	CameraMake        string
	CameraType        string
	CameraOwner       string
	CameraDescription string `gorm:"type:text;"`
	CameraNotes       string `gorm:"type:text;"`
}

func NewCamera(modelName string, makeName string) *Camera {
	if modelName == "" {
		modelName = "Unknown"
	}

	var cameraSlug string

	if makeName != "" {
		cameraSlug = slug.MakeLang(makeName+" "+modelName, "en")
	} else {
		cameraSlug = slug.MakeLang(modelName, "en")
	}

	result := &Camera{
		CameraModel: modelName,
		CameraMake:  makeName,
		CameraSlug:  cameraSlug,
	}

	return result
}

func (c *Camera) FirstOrCreate(db *gorm.DB) *Camera {
	db.FirstOrCreate(c, "camera_model = ? AND camera_make = ?", c.CameraModel, c.CameraMake)

	return c
}

func (c *Camera) String() string {
	if c.CameraMake != "" && c.CameraModel != "" {
		return fmt.Sprintf("%s %s", c.CameraMake, c.CameraModel)
	} else if c.CameraModel != "" {
		return c.CameraModel
	}

	return ""
}
