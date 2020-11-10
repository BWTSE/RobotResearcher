package router

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func (r *Router) applyCodeRoutes(rg *gin.RouterGroup) {
	rg.Use(r.authMiddleware)

	rg.POST("/run", func(c *gin.Context) {
		code, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = os.Remove("tmp.java")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = ioutil.WriteFile("tmp.java", []byte(code), 0644)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		out, err := exec.Command("/usr/bin/java", "-Djava.security.manager", "./tmp.java").Output()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"out": string(out),
		})
	})
}
