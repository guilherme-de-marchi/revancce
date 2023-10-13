package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/guilherme-de-marchi/revancce/api/v1/service"
)

func (c Controllers) PaymentWebhookOpenpix() {
	c.Group.POST("/payment/webhook/openpix", paymentWebhookOpenpix)
}

func paymentWebhookOpenpix(c *gin.Context) {
	// var buf bytes.Buffer
	// tee := io.TeeReader(c.Request.Body, &buf)
	// b, _ := io.ReadAll(tee)
	// println(string(b))

	// h := hmac.New(sha1.New, []byte("-----------"))
	// h.Write(b)

	// if c.GetHeader("X-OpenPix-Signature") != base64.StdEncoding.EncodeToString(h.Sum(nil)) {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	var req model.PaymentWebhookOpenpixPostReq
	if err := c.ShouldBind(&req); err != nil {
		pkg.Log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(req.Charge.AdditionalInfo) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var batch string
	for _, v := range req.Charge.AdditionalInfo {
		k, ok := v["key"]
		if !ok || k != "batch-id" {
			continue
		}

		k, ok = v["value"]
		if !ok {
			continue
		}

		if k == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		batch = k
		break
	}

	c.JSON(service.PaymentWebhookOpenpix(c, model.PaymentWebhookOpenpixPostIn{
		Type:        req.Event,
		Name:        req.Charge.Payer.Name,
		CPF:         req.Charge.Payer.TaxID.TaxID,
		Email:       req.Charge.Payer.Email,
		Phone:       req.Charge.Payer.Phone,
		Transaction: req.Charge.CorrelationID,
		Batch:       batch,
	}))
}
