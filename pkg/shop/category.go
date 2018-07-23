package shop

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model `structs:",flatten"`
	Name       string    `validate:"required"`
	Parent     *Category `gorm:"foreignkey:ParentID"`
	ParentID   uint
}
