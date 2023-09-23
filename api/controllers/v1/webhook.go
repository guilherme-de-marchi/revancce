package v1

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/charge"
	wh "github.com/stripe/stripe-go/v75/webhook"
)

func (g group) Webhook() {
	g.Group.POST("/webhook", webhook)
}

func webhook(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, pkg.Error(err))
		return
	}

	event, err := wh.ConstructEvent(
		body,
		c.GetHeader("Stripe-Signature"),
		pkg.StripeWebhookSecret,
	)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, pkg.Error(err))
		return
	}

	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			c.AbortWithError(http.StatusInternalServerError, pkg.Error(err))
			return
		}

		if paymentIntent.LatestCharge == nil {
			c.AbortWithError(http.StatusInternalServerError, pkg.Error(pkg.ErrWebhookFieldMalformatted))
			return
		}

		latestCharge, err := charge.Get(paymentIntent.LatestCharge.ID, nil)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, pkg.Error(err))
			return
		}

		if latestCharge.BillingDetails == nil {
			c.AbortWithError(http.StatusInternalServerError, pkg.Error(pkg.ErrWebhookFieldMalformatted))
			return
		}

		message := mail.NewSingleEmail(
			mail.NewEmail(
				pkg.SendGridPurchasesName,
				pkg.SendGridPurchasesEmail,
			),
			"Assunto legal com o mano "+latestCharge.BillingDetails.Name,
			mail.NewEmail(
				latestCharge.BillingDetails.Name,
				latestCharge.BillingDetails.Email,
			),
			"texto maneiro para o piá "+latestCharge.BillingDetails.Name,
			"<h1>HTML legalzinho</h1>",
		)

		_, err = pkg.Mail.Send(message)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, pkg.Error(err))
			return
		}

	}
}
