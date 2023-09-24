package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api"
)

func main() {
	e := gin.Default()
	e.Static("/static", "./web/public")

	if err := api.Setup(e); err != nil {
		log.Fatal(err)
	}

	e.Run(":8080")
}
