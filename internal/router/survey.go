package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sidus.io/robotresearcher/internal/database"
)

func (r *Router) applySurveyRoutes(rg *gin.RouterGroup) {
	rg.Use(r.authMiddleware)
	rg.Use(r.agreementMiddleware)
	rg.Use(r.scenarioMiddleware)

	rg.POST("/", func(c *gin.Context) {
		s := getSession(c)

		if s.SurveyAnswers != nil {
			c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
				"error": "already submitted",
			})

			return
		}

		var submission database.SurveySubmission

		err := c.BindJSON(&submission)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		err = r.database.SaveSurveyAnswers(s.ID, submission)
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