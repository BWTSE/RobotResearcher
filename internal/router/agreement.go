package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) applyAgreementRoutes(rg *gin.RouterGroup) {
	rg.Use(r.authMiddleware)
	rg.POST("/accept", func(c *gin.Context) {
		s := getSession(c)

		err := r.database.AcceptAgreement(s.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
}

func (r *Router) agreementMiddleware(c *gin.Context) {
	s := getSession(c)
	if s.AgreementAccepted {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
			"error": "agreement was not accepted",
		})
	}
}
