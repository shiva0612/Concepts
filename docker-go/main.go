package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	log.Println("starting gin server...")
	log.Fatalln(r.Run(":8080"))
}
