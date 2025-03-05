package handlers

import (
	"library-management1/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	user, _ := c.Get("currentUser")
	var userData models.User
	userData = user.(models.User)

	// abhi or karna hai

}
