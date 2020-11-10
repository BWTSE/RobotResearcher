package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"sidus.io/robotresearcher/internal/database"
)

func (r *Router) authMiddleware(c *gin.Context) {
	id := c.GetHeader("Authorization")
	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no id token",
		})

		return
	}

	session, err := r.database.GetSession(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid id token",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		return
	}

	if session.EndedAt.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "session ended",
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
		id, err := r.database.CreateSession(database.Session{
			RegisterCode: "",  //todo
			Scenarios:    nil, //todo
			StartedAt:    time.Now(),
			EndedAt:      time.Now().Add(2 * time.Hour), //TODO
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.Header("Authorization", id)
		c.JSON(http.StatusOK, gin.H{})
	})

	rg.POST("/close", r.authMiddleware, func(c *gin.Context) {
		c.Header("Authorization", "")
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	rg.GET("/status", r.authMiddleware, func(c *gin.Context) {
		s := getSession(c)
		c.JSON(http.StatusOK, gin.H{
			"id": s.ID.Hex(),
		})
	})
}
