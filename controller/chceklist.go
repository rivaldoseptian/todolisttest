package controllers

import (
	"net/http"
	"todolist-api/config"
	"todolist-api/models"
	"todolist-api/utils"

	"github.com/gin-gonic/gin"
)

func GetChecklists(c *gin.Context) {
	var checklists []models.Checklist
	config.DB.Find(&checklists)
	c.JSON(http.StatusOK, gin.H{"data": checklists})
}

func CreateChecklist(c *gin.Context) {
	var checklist models.Checklist
	if err := c.ShouldBindJSON(&checklist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	checklist.UserID = userID.(uint)
	config.DB.Create(&checklist)

	c.JSON(http.StatusCreated, gin.H{"message": "Checklist created", "data": checklist})
}
func DeleteChecklist(c *gin.Context) {
	id := c.Param("checklistId")
	if err := config.DB.Delete(&models.Checklist{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete checklist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checklist deleted successfully"})
}
func GetChecklistItems(c *gin.Context) {
	checklistID := c.Param("checklistId")
	var items []models.ChecklistItem

	if err := config.DB.Where("checklist_id = ?", checklistID).Find(&items).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checklist not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}
func CreateChecklistItem(c *gin.Context) {
	var item models.ChecklistItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ChecklistID = utils.ParseUintParam(c.Param("checklistId"))
	config.DB.Create(&item)

	c.JSON(http.StatusCreated, gin.H{"message": "Item added to checklist", "data": item})
}
func GetChecklistItemByID(c *gin.Context) {
	var item models.ChecklistItem
	checklistID := c.Param("checklistId")
	itemID := c.Param("checklistItemId")

	if err := config.DB.Where("id = ? AND checklist_id = ?", itemID, checklistID).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}
func UpdateChecklistItemStatus(c *gin.Context) {
	var item models.ChecklistItem
	itemID := c.Param("checklistItemId")

	if err := config.DB.First(&item, itemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	item.Status = !item.Status
	config.DB.Save(&item)

	c.JSON(http.StatusOK, gin.H{"message": "Item status updated", "data": item})
}
func DeleteChecklistItem(c *gin.Context) {
	itemID := c.Param("checklistItemId")

	if err := config.DB.Delete(&models.ChecklistItem{}, itemID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
func RenameChecklistItem(c *gin.Context) {
	var item models.ChecklistItem
	itemID := c.Param("checklistItemId")

	if err := config.DB.First(&item, itemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item renamed", "data": item})
}
