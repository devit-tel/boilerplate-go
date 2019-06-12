// Package server provides functionality for handling data input and queries
package server

import (
	"context"
	"net/http"

	"github.com/NV4RE/boilerplate-go/pkg/commons"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// REST provides functionality for HTTP REST API Server
type REST struct {
	server *http.Server
}

// CreateServer create a server
func CreateServer(options commons.Options) *REST {
	router := gin.Default()

	// v1 routers
	v1 := router.Group("/v1")
	{
		v1.GET("/login/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

	return &REST{
		server: server,
	}

}

// Stop stops REST API gracefully
func (r *REST) Stop() {
	logrus.Warn("Stopping REST server..")
	r.server.Shutdown(context.TODO())
}
