package routes

import (
	controllers "todolist-api/controller"
	"todolist-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Grup utama dengan prefix /api
	api := r.Group("/api")

	// Rute Autentikasi (tanpa middleware)
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// Grup rute yang memerlukan autentikasi JWT
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/checklist", controllers.GetChecklists)                   // Get All Checklist
		auth.POST("/checklist", controllers.CreateChecklist)                // Create Checklist
		auth.DELETE("/checklist/:checklistId", controllers.DeleteChecklist) // Delete Checklist

		auth.GET("/checklist/:checklistId/item", controllers.GetChecklistItems)                           // Get Checklist Items
		auth.POST("/checklist/:checklistId/item", controllers.CreateChecklistItem)                        // Create Checklist Item
		auth.GET("/checklist/:checklistId/item/:checklistItemId", controllers.GetChecklistItemByID)       // Get Checklist Item
		auth.PUT("/checklist/:checklistId/item/:checklistItemId", controllers.UpdateChecklistItemStatus)  // Update Item Status
		auth.DELETE("/checklist/:checklistId/item/:checklistItemId", controllers.DeleteChecklistItem)     // Delete Item
		auth.PUT("/checklist/:checklistId/item/rename/:checklistItemId", controllers.RenameChecklistItem) // Rename Item
	}

}
