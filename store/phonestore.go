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

func (ps *PhoneStore) GetAll() ([]PhoneNumber, error) {
	var phoneNums []PhoneNumber
	result := ps.db.Find(&phoneNums)
	if err := result.Error; err != nil {
		return nil, err
	}
	return phoneNums, nil
}

// Takes as input the function to apply to each DB entry to normalize it
func (ps *PhoneStore) Normalize(fn func(string) (string, error)) error {
	phoneNums, err := ps.GetAll()
	if err != nil {
		return err
	}

	for _, pn := range phoneNums {
		pn.Number, err = fn(pn.Number)
		if err != nil {
			return err
		}
		result := ps.db.Save(&pn)
		if err := result.Error; err != nil {
			return err
		}
	}

	return nil
}

func (ps *PhoneStore) RemoveDupes() error {
	phoneNums, err := ps.GetAll()
	if err != nil {
		return err
	}

	dupesMap := make(map[string][]uint)
	for _, pn := range phoneNums {
		ids := dupesMap[pn.Number]
		dupesMap[pn.Number] = append(ids, pn.ID)
	}

	for _, ids := range dupesMap {
		if len(ids) == 1 {
			continue
		}
		result := ps.db.Delete(&phoneNums, ids[1:])
		if err := result.Error; err != nil {
			return err
		}
	}
	return nil
}
