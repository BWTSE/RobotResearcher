package router

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *Router) applyCodeRoutes(rg *gin.RouterGroup) {
	rg.Use(r.authMiddleware)
	rg.Use(r.agreementMiddleware)

	rg.POST("/run", func(c *gin.Context) {
		code, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = ioutil.WriteFile("ttt/main.java", []byte(code), 0644)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer os.Remove("tt/main.java")

		cd, cancel := context.WithTimeout(c, time.Second*10)
		defer cancel()

		// TODO: policy is not working
		cmd := exec.CommandContext(cd, "/usr/bin/java" /*,"-Djava.security.manager" "-Djava.security.debug=all", "-Djava.security.policy==./untrusted.policy",*/, "-Xmx10M", "./main.java")

		out, err := cmd.CombinedOutput()
		if err != nil {
			_, ok := err.(*exec.ExitError)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(200, gin.H{
			"out":  string(out),
			"code": cmd.ProcessState.ExitCode(),
		})
	})
}
