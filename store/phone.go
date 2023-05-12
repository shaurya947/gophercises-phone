package store

import "gorm.io/gorm"

type PhoneNumber struct {
	gorm.Model
	Number string
}
