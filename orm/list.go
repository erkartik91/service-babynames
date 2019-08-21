package orm

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type List struct {
	ID        string `gorm:"primary_key"`
	Name      string
	BabyNames pq.StringArray `gorm:"type:varchar(100)[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (l *List) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

type ListORM struct {
	DB *gorm.DB
}

func (l *ListORM) Create(name string) (*List, error) {
	if name == "" {
		return nil, errors.New("Missing list name")
	}

	list := List{
		Name:      name,
		BabyNames: pq.StringArray{},
	}
	return &list, l.DB.Save(&list).Error
}

func (l *ListORM) Update(list *List) (*List, error) {
	return list, l.DB.Update(list).Error
}

func (l *ListORM) Read(id string) (*List, error) {
	list := List{
		ID: id,
	}

	return &list, l.DB.Where(&list).Find(&list).Error
}
