package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
)

type Group struct {
	Group *gin.RouterGroup
}

func (g Group) GetProduct() {
	g.Group.GET("/product/:id", getProduct)
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":    c.Param("id"),
		"price": 20.0,
	})
}

func (g Group) CreateCheckoutSession() {
	g.Group.POST("/product/:id/create-checkout-session", createCheckoutSession)
}

func createCheckoutSession(c *gin.Context) {
	domain := "http://localhost:8080"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_1NsyLqJv8u5au94NyNzKF3sr"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/static/success.html"),
		CancelURL:  stripe.String(domain + "/static/cancel.html"),
	}

	s, err := session.New(params)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusSeeOther, s.URL)
}
