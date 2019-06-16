// Package http provides functionality for handling data input and queries
package server

import (
	"context"
	"github.com/NV4RE/boilerplate-go/pkg/http/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// REST provides functionality for HTTP REST API Server
type HttpServer struct {
	server *http.Server
}

// CreateServer create a http
func CreateServer(addr string) *HttpServer {
	r := gin.Default()

	// v1 routers
	v1 := r.Group("/v1")
	{
		sample := v1.Group("/sample")
		{
			sample.GET("/:name", router.SampleParam)
			sample.POST("/", router.SampleJsonBody)
		}
	}

	return &HttpServer{
		server: &http.Server{
			Addr:    addr,
			Handler: r,
		},
	}
}

func (r *HttpServer) Listen() {
	logrus.Fatal(r.server.ListenAndServe())
}

func (r *HttpServer) Stop() {
	logrus.Info("Stopping REST http..")
	logrus.Fatal(r.server.Shutdown(context.TODO()))
}
