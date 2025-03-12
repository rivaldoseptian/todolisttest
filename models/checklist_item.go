package models

type ChecklistItem struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ChecklistID uint   `json:"checklist_id"`
	ItemName    string `json:"itemName" binding:"required"`
	Status      bool   `json:"status"`
}
