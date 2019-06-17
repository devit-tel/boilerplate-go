package router

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


func Health(c *gin.Context) {
    // TODO add real health check
	c.String(http.StatusOK, "OK")
}

