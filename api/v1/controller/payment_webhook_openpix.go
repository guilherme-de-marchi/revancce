package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/guilherme-de-marchi/revancce/api/v1/model"
	"github.com/jackc/pgx/v5"
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

	var client string
	row := pkg.Database.QueryRow(
		c,
		`
			select id
			from clients
			where cpf=$1
		`,
		req.Pix.Payer.TaxID.TaxID,
	)
	err := row.Scan(&client)
	if err != nil && err != pgx.ErrNoRows {
		pkg.Log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err == pgx.ErrNoRows {
		row = pkg.Database.QueryRow(
			c,
			`
				insert into clients
				(name, email, cpf, phone)
				values ($1, $2, $3, $4)
				returning id
			`,
			req.Pix.Payer.Name,
			req.Pix.Payer.Email,
			req.Pix.Payer.TaxID.TaxID,
			req.Pix.Payer.Phone,
		)
		if err = row.Scan(&client); err != nil {
			pkg.Log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	_, err = pkg.Database.Exec(
		c,
		`
			insert into clients_tickets
			(client, batch, transaction)
			values ($1, $2, $3)
		`,
		client,
		batch,
		req.Charge.TransactionID,
	)
	if err != nil {
		pkg.Log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
