package router

import (
	"github.com/gildemberg-santos/libratica-go/internal/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Route *gin.Engine
}

func NewRoute() *Router {
	return &Router{
		Route: gin.Default(),
	}
}

func (r *Router) setting() {
	r.Route.Use(gin.Recovery())
	r.Route.Use(controller.AuthMiddleware)
}

func (r *Router) Run() {
	r.Route.POST("/login", controller.Auth)
	r.setting()
	// TODO: Implement the other routes
	r.Route.Run()
}
