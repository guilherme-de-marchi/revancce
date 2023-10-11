package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.Static("/static", "./web/public")

	if err := Setup(e); err != nil {
		log.Fatal(err)
	}

	e.Run(":8080")
}
