package models

type Checklist struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"itemName" binding:"required"`
	UserID uint   `json:"user_id"`
}
