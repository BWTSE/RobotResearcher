package router

import (
	"github.com/gin-gonic/gin"

	"sidus.io/robotresearcher/internal/database"

	"sidus.io/robotresearcher/internal/code"
)

type Router struct {
	codeService *code.CodeService
	database    *database.Database
}

func NewRouter(codeService *code.CodeService, database *database.Database) (*Router, error) {
	return &Router{
		codeService: codeService,
		database:    database,
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

	api := e.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.applyCodeRoutes(api.Group("/code"))
	r.applyAuthRoutes(api.Group("/auth"))
	r.applyAgreementRoutes(api.Group("/agreement"))

	return e.Run("0.0.0.0:" + port)
}
