package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PhoneStore struct {
	db *gorm.DB
}

func NewStore(url string) (*PhoneStore, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&PhoneNumber{})
	return &PhoneStore{db}, nil
}

func (ps *PhoneStore) Insert(phoneNums []*PhoneNumber) error {
	result := ps.db.Create(phoneNums)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected != int64(len(phoneNums)) {
		return fmt.Errorf(
			"Tried to insert %d records but only inserted %d",
			len(phoneNums),
			result.RowsAffected,
		)
	}

	return nil
}
