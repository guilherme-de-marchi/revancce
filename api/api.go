package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	v1 "github.com/guilherme-de-marchi/revancce/api/v1"
	"github.com/redis/go-redis/v9"
	"github.com/sendgrid/sendgrid-go"
	"github.com/stripe/stripe-go/v75"
)

func Setup(e *gin.Engine) error {
	ctx := context.Background()

	stripe.Key = pkg.StripeSecretKey

	pkg.Memory = redis.NewClient(&redis.Options{
		Addr:     pkg.RedisAddr,
		Password: pkg.RedisPassword,
	})
	if err := pkg.Memory.Ping(ctx).Err(); err != nil {
		return err
	}

	pkg.Mail = sendgrid.NewSendClient(pkg.SendGridSecretKey)

	v1.Setup(e)

	return nil
}
