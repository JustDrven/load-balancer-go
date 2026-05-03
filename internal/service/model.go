package service

type Service struct {
	ID            uint `gorm:"primarykey"`
	Type          string
	SSL           bool
	Address       string
	MaxReferences int
}

type ManagedService struct {
	Address       string
	MaxReferences int
	References    int
}
