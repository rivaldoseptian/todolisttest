package models

type ChecklistItem struct {
	ID          uint `gorm:"primaryKey"`
	ChecklistID uint
}
