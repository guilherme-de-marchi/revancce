package controllers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	wh "github.com/stripe/stripe-go/v75/webhook"
)

var SENDGRID_SECRET_KEY = os.Getenv("SENDGRID_SECRET_KEY")

func (g Group) Webhook() {
	g.Group.POST("/webhook", webhook)
}

func webhook(c *gin.Context) {
	var body any
	if err := c.Bind(&body); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	endpointSecret := "whsec_efddb06c656923f000eb3ef8e80fc5813adfbfc9e4446b3c2a6ba87f6994d7c1"
	_, err = wh.ConstructEvent(rawBody, c.GetHeader("Stripe-Signature"), endpointSecret)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	message := mail.NewSingleEmail(
		mail.NewEmail("Vendas", "flor@revancce.com"),
		"Assunto legal",
		mail.NewEmail("Cliente", "demarchi@revancce.com"),
		"texto maneiro",
		"<h1>HTML legalzinho</h1>",
	)

	client := sendgrid.NewSendClient(SENDGRID_SECRET_KEY)
	response, err := client.Send(message)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	log.Println(response)
	// log.Println(body)
}
