package models

type Checklist struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
