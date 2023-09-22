package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/controllers"
	"github.com/stripe/stripe-go/v75"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	r := gin.Default()
	r.Static("/static", "./web/public")

	g := controllers.Group{Group: r.Group("/api")}
	g.GetProduct()
	g.CreateCheckoutSession()
	g.Webhook()

	r.Run(":8080")
}
