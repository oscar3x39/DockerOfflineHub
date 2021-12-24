package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
)

func requestHandle(c *gin.Context) {
	var m map[string]interface{}
	err := c.Bind(&m)
	if err != nil {
		return
	}

	if m["name"] == nil {
		handler(c, "false")
		return
	}

	name := fmt.Sprintf("%v", m["name"])
	re := regexp.MustCompile(`^[a-z0-9\.\/]*$`).FindString(name)
	if name != re {
		handler(c, "false")
		return
	}

	fmt.Println("pulling docker image from " + name)
	if dockerPullImage(name) {
		if dockerSaveImage(name) {
			handler(c, "true")
		}
	}

	handler(c, "false")
}

func commandExecute(args ...string) bool {
	cmd := exec.Command("docker", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	fmt.Println("Result: " + out.String())
	return true
}

func dockerSaveImage(name string) bool {
	cmd := []string{"save", name, "../public/" + name}
	return commandExecute(cmd...)
}

func dockerPullImage(name string) bool {
	cmd := []string{"pull", name}
	return commandExecute(cmd...)
}

func handler(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"success": msg,
	})
}

func main() {
	r := gin.Default()
	r.POST("/pull", requestHandle)
	r.Run(":80")
}
