package shop

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model  `structs:",flatten"`
	Name        string `validate:"required"`
	Price       uint   `validate:"required"`
	Description string `validate:"required"`
	Available   bool
	Stock       uint
	CategoryID  int
	Category    Category `validate:"-"`
}
