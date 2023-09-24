package repository

import (
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
)

func GetProduct(id string) any {
	return map[string]any{
		"id":    id,
		"price": 20.0,
	}
}

func PurchaseProduct() (*stripe.CheckoutSession, error) {
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
	return s, pkg.Error(err)
}
