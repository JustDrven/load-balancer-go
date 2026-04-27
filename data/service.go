package data

import "gorm.io/gorm"

type ServiceStatus int

const (
	Success ServiceStatus = iota
	Failed
)

type Service struct {
	gorm.Model

	Type          string
	Address       string
	MaxReferences int
}

type ManagedService struct {
	Address       string
	MaxReferences int
	References    int
}
