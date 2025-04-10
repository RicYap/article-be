package article

import (
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
}

// New ...
func New(db *gorm.DB) *Data {
	return &Data{
		db: db,
	}
}
