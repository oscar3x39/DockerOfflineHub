package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	var m map[string]interface{}
	err := c.Bind(&m)
	if err != nil {
		return
	}

	// cmd := exec.Command("docker", "pull", "")
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("%v\n", m)
	c.JSON(200, gin.H{
		"success": true,
	})
}

func main() {
	r := gin.Default()
	r.POST("/cmd", handler)
	r.Run()
}
