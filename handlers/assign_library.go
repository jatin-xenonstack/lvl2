package handlers

import (
	"library-management1/database"
	"library-management1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssignAdmin(c *gin.Context) {
	user, err := c.Get("currentUser")
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no user found"})
		return
	}

	var userData models.User
	userData = user.(models.User)
	if userData.Role == "Owner" {
		var admin models.Admin
		if err := c.ShouldBindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var changeRole models.User
		database.DB.Where("id = ?", admin.ID).Find(&changeRole)
		if changeRole.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User do not exist"})
			return
		}
		database.DB.Model(models.User{}).Where("id = ?", admin.ID).Update("Role", "admin")
		var library models.Library
		database.DB.Where("id = ?", userData.ID).Find(&library)
		libraryData := models.LibraryUser{
			UserId:    admin.ID,
			LibraryId: library.ID,
		}
		database.DB.Create(&libraryData)
		c.JSON(http.StatusOK, gin.H{"message": "Admin Assigned Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request Rejected"})
	}
}
