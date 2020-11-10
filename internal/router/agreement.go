package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) applyAgreementRoutes(rg *gin.RouterGroup) {
	rg.Use(r.authMiddleware)
	rg.POST("/accept", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}
