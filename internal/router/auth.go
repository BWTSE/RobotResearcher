package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) applyAuthRoutes(rg *gin.RouterGroup) {
	rg.GET("/status", func(c *gin.Context) {
		id := c.GetHeader("Authorization")
		if id == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "no id token",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	rg.POST("/register", func(c *gin.Context) {
		c.Header("Authorization", "12345")
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	rg.POST("/close", func(c *gin.Context) {
		c.Header("Authorization", "")
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}
