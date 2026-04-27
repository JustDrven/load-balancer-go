package data

import "gorm.io/gorm"

type Service struct {
	gorm.Model

	Type          string
	Address       string
	MaxReferences int
}
