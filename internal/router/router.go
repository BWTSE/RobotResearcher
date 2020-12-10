package router

import (
	"github.com/gin-gonic/gin"

	"sidus.io/robotresearcher/internal/database"

	"sidus.io/robotresearcher/internal/code"
)

type Router struct {
	codeService *code.CodeService
	database    *database.Database
	tempDir     string
}

func NewRouter(codeService *code.CodeService, database *database.Database, tempDir string) (*Router, error) {
	return &Router{
		codeService: codeService,
		database:    database,
		tempDir:     tempDir,
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
	r.applyScenarioRoutes(api.Group("/scenario"))
	r.applySurveyRoutes(api.Group("/survey"))

	return e.Run("0.0.0.0:" + port)
}