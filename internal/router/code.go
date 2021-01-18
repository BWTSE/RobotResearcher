package router

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Router) applyCodeRoutes(rg *gin.RouterGroup) {
	rg.Use(r.authMiddleware)
	rg.Use(r.agreementMiddleware)

	rg.POST("/run", func(c *gin.Context) {
		var request struct {
			Submission map[string]string
		}

		err := c.BindJSON(&request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		s := getSession(c)

		packageName := "unknown"
		for i := len(s.Scenarios) - 1; i >= 0; i-- {
			if !s.Scenarios[i].StartedAt.IsZero() {
				packageName = s.Scenarios[i].Name
				break
			}
		}

		dir := path.Join(r.tempDir, "robotresearcher", uuid.New().String())

		err = os.MkdirAll(path.Join(dir, packageName), 0775)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer os.RemoveAll(dir)

		for name, content := range request.Submission {
			if strings.Contains(name, string(os.PathSeparator)) {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no hacking"})
				return
			}

			if !strings.HasSuffix(name, ".java") {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "only java files"})
				return
			}

			if content == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file must have content"})
				return
			}

			code, err := base64.StdEncoding.DecodeString(content)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			err = ioutil.WriteFile(path.Join(dir, packageName, name), code, 0644)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		for name := range request.Submission {
			cmd := exec.Command(
				"javac",
				path.Join(packageName, name),
			)
			cmd.Dir = dir

			out, err := cmd.CombinedOutput()
			if err != nil {
				_, ok := err.(*exec.ExitError)
				if ok {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": "compilation failed",
						"out":   string(out),
						"code":  cmd.ProcessState.ExitCode(),
					})
					return
				} else {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			}
		}

		cd, cancel := context.WithTimeout(c, time.Second*10)
		defer cancel()

		cmd := exec.CommandContext(
			cd,
			"java",
			"-Djava.security.manager",
			//			"-Djava.security.debug=all",
			//			"-Djava.security.policy==./../../untrusted.policy",
			"-Xmx10M",
			packageName+".Main",
		)
		cmd.Dir = dir

		out, err := cmd.CombinedOutput()
		if err != nil {
			_, ok := err.(*exec.ExitError)
			if !ok {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(200, gin.H{
			"out":  string(out),
			"code": cmd.ProcessState.ExitCode(),
		})
	})
}
