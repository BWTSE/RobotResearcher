package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"sidus.io/robotresearcher/internal/database"
)

func (r *Router) authMiddleware(c *gin.Context) {
	id := c.GetHeader("Authorization")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "no id token",
		})

		return
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid id hex",
		})

		return
	}

	session, err := r.database.GetSession(oid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid id token",
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		return
	}

	if !session.EndedAt.IsZero() {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "session canceled",
		})

		return
	}

	if session.StartedAt.Before(time.Now().Add(-10 * time.Hour)) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "session expired",
		})

		return
	}

	c.Set("session", session)
	c.Next()
}

func getSession(c *gin.Context) database.Session {
	session, _ := c.Get("session")
	return session.(database.Session)
}

func (r *Router) applyAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", func(c *gin.Context) {
		var request struct {
			Code string
		}

		err := c.BindJSON(&request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		valid, ignoreCount, err := r.database.IsValidCode(request.Code)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}
		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Wrong Code",
			})
			return
		}

		var sequence []database.Scenario
		if ignoreCount {
			sequence = r.scenarioService.GetDummySequence()
		} else {
			sequence = r.scenarioService.GetSequence()
		}

		id, err := r.database.CreateSession(request.Code, sequence, ignoreCount)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}
		c.Header("Authorization", id.Hex())
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	rg.POST("/close", r.authMiddleware, func(c *gin.Context) {
		err := r.database.EndSession(getSession(c).ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.Header("Authorization", "")
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	rg.GET("/status", r.authMiddleware, func(c *gin.Context) { // todo
		s := getSession(c)

		var scenarioNames []string
		for _, scenario := range s.Scenarios {
			scenarioNames = append(scenarioNames, scenario.Name)
		}

		c.JSON(http.StatusOK, gin.H{
			"id":             s.ID.Hex(),
			"scenario_order": scenarioNames,
		})
	})
}
