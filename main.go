package main

import "gopkg.in/gin-gonic/gin.v1"
import "net/http"

func main() {
  router := gin.Default()

  router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Health OK")
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello le mec qui dors au premier rang%s", name)
	})

  router.Run(":80")
}
