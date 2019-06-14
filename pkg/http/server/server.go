// Package http provides functionality for handling data input and queries
package server

import (
	"context"
	"net/http"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// REST provides functionality for HTTP REST API Server
type HttpServer struct {
	server *http.Server
}

// CreateServer create a http
func CreateServer(addr string) *HttpServer {
	router := gin.Default()

	// v1 routers
	v1 := router.Group("/v1")
	{
		v1.GET("/login/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
	}

	return &HttpServer{
		server: &http.Server{
            Addr:    addr,
            Handler: router,
        },
	}
}

func(r *HttpServer) Start() {
    r.server.ListenAndServe()
}

// Stop stops REST API gracefully
func (r *HttpServer) Stop() {
	logrus.Warn("Stopping REST http..")
	r.server.Shutdown(context.TODO())
}
