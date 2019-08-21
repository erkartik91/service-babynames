package orm

import (
	"github.com/jinzhu/gorm"
)

type ORM struct {
	List *ListORM
}

func autoMigrate(DB *gorm.DB) *gorm.DB {
	return DB.
		AutoMigrate(&List{})
}

func New(DB *gorm.DB) *ORM {
	// Uncomment to see SQL queries in test debug logs
	// DB.LogMode(true)

	db := autoMigrate(DB)

	return &ORM{
		List: &ListORM{
			DB: db,
		},
	}
}
