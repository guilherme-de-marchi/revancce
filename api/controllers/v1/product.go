package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
)

func (g group) GetProduct() {
	g.Group.GET("/product/:id", getProduct)
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":    c.Param("id"),
		"price": 20.0,
	})
}

func (g group) Purchase() {
	g.Group.POST("/product/:id/purchase", purchase)
}

func purchase(c *gin.Context) {
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_1NsyLqJv8u5au94NyNzKF3sr"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(pkg.ServerDomain + pkg.ProductPurchaseRedirectSuccess),
		CancelURL:  stripe.String(pkg.ServerDomain + pkg.ProductPurchaseRedirectCancel),
	}

	s, err := session.New(params)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, pkg.Error(err))
		return
	}

	c.Redirect(http.StatusSeeOther, s.URL)
}
