package repository

import (
	"encoding/json"

	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/charge"
	wh "github.com/stripe/stripe-go/v75/webhook"
)

func Webhook(body []byte, stripeSignatureHeader string) error {
	event, err := wh.ConstructEvent(
		body,
		stripeSignatureHeader,
		pkg.StripeWebhookSecret,
	)
	if err != nil {
		return pkg.Error(err)
	}

	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
			return pkg.Error(err)
		}

		if paymentIntent.LatestCharge == nil {
			return pkg.Error(pkg.ErrWebhookFieldMalformatted)
		}

		latestCharge, err := charge.Get(paymentIntent.LatestCharge.ID, nil)
		if err != nil {
			return pkg.Error(err)
		}

		if latestCharge.BillingDetails == nil {
			return pkg.Error(pkg.ErrWebhookFieldMalformatted)
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
			return pkg.Error(err)
		}
	}

	return nil
}
