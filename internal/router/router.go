package router

import (
	"github.com/gin-gonic/gin"

	"sidus.io/robotresearcher/internal/code"
)

type Router struct {
	codeService *code.CodeService
}

func NewRouter(codeService *code.CodeService) (*Router, error) {
	return &Router{
		codeService: codeService,
	}, nil
}

func (r *Router) Serve(debug bool, port string) error {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.applyCodeRoutes(e.Group("code"))

	return e.Run("0.0.0.0:" + port)
}