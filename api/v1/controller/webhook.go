package controller

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/repository"
)

func (c Controllers) Webhook() {
	c.Group.POST("/webhook", webhook)
}

func webhook(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, pkg.Error(err))
		return
	}

	stripeSignatureHeader := c.GetHeader("Stripe-Signature")
	if stripeSignatureHeader == "" {
		c.AbortWithError(http.StatusBadRequest, pkg.Error(errors.New("'Stripe-Signature' header is missing")))
		return
	}

	if err := repository.Webhook(body, stripeSignatureHeader); err != nil {
		c.AbortWithError(http.StatusInternalServerError, pkg.Error(err))
		return
	}
}
