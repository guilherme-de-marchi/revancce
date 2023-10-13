package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
)

func main() {
	e := gin.Default()
	e.Use(pkg.Cors())
	// e.Use(cors.Default())

	if err := Setup(e); err != nil {
		log.Fatal(err)
	}

	e.Run(":8080")
}
