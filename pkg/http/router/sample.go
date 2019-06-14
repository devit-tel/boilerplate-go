package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func SampleParam(c *gin.Context) {
	name := c.Param("name")
	username := c.DefaultQuery("username", "Guest")
	c.String(http.StatusOK, "Hello %s (%s)", name, username)
}

func SampleJsonBody(c *gin.Context) {
	var json Login
	// If not error (Valid JSON)
	if c.BindJSON(&json) == nil {
		c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("you are logged in as %s:%s", json.Username, json.Password)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Not valid JSON"})
	}
}
