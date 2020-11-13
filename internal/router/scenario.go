package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *Router) scenarioMiddleware(c *gin.Context) {
	//todo: check that all scenarios are completed

}

// todo, no late submissions?
func (r *Router) applyScenarioRoutes(rg *gin.RouterGroup) {
	rg.Use(r.authMiddleware)
	rg.Use(r.agreementMiddleware)

	rg.GET("/:num", func(c *gin.Context) {
		s := getSession(c)

		n, err := strconv.Atoi(c.Param("num"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "scenario parameter is not a number",
			})

			return
		}

		if n < 0 || len(s.Scenarios) <= n {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "no such scenario",
			})

			return
		}

		for i := n - 1; i >= 0; i-- {
			if s.Scenarios[i].SubmittedAt.IsZero() {
				c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
					"error": "previous scenario not completed",
				})

				return
			}
		}

		scenario := s.Scenarios[n]

		if !scenario.SubmittedAt.IsZero() {
			c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
				"error": "scenario is already submitted",
			})

			return
		}

		if scenario.StartedAt.IsZero() {
			err = r.database.OpenScenario(s.ID, n)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})

				return
			}
			scenario.StartedAt = time.Now()
		}

		c.JSON(http.StatusOK, gin.H{
			"info":       scenario.Instructions,
			"started_at": scenario.StartedAt,
			"given":      scenario.Starting,
		})
	})

	rg.POST("/:num", func(c *gin.Context) {
		s := getSession(c)

		n, err := strconv.Atoi(c.Param("num"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "scenario parameter is not a number",
			})

			return
		}

		if n < 0 || len(s.Scenarios) <= n {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "no such scenario",
			})

			return
		}

		for i := n - 1; i >= 0; i-- {
			if s.Scenarios[i].SubmittedAt.IsZero() {
				c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
					"error": "previous scenario not completed",
				})

				return
			}
		}

		scenario := s.Scenarios[n]

		if !scenario.SubmittedAt.IsZero() {
			c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
				"error": "scenario is already submitted",
			})

			return
		}

		var request struct {
			Submission map[string]string
		}

		err = c.BindJSON(&request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		err = r.database.CommitScenario(s.ID, n, request.Submission)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "done",
		})
	})
}
