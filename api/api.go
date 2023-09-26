package api

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	v1 "github.com/guilherme-de-marchi/revancce/api/v1"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"github.com/sendgrid/sendgrid-go"
	"github.com/stripe/stripe-go/v75"
)

func Setup(e *gin.Engine) error {
	ctx := context.Background()

	stripe.Key = pkg.StripeSecretKey

	pkg.Log = log.New(
		os.Stdout,
		"",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Llongfile,
	)

	var err error
	pkg.Database, err = pgx.Connect(ctx, pkg.DatabaseURL)
	if err != nil {
		return err
	}
	if err := pkg.Database.Ping(ctx); err != nil {
		return err
	}

	pkg.Memory = redis.NewClient(&redis.Options{
		Addr:     pkg.MemoryAddr,
		Password: pkg.MemoryPassword,
	})
	if err := pkg.Memory.Ping(ctx).Err(); err != nil {
		return err
	}

	pkg.Mail = sendgrid.NewSendClient(pkg.SendGridSecretKey)

	pkg.RSA, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	v1.Setup(e)

	return nil
}
