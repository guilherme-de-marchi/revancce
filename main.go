package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	v1 "github.com/guilherme-de-marchi/revancce/api/controllers/v1"
	"github.com/guilherme-de-marchi/revancce/api/pkg"
	"github.com/redis/go-redis/v9"
	"github.com/sendgrid/sendgrid-go"
	"github.com/stripe/stripe-go/v75"
)

func main() {
	ctx := context.Background()

	stripe.Key = pkg.StripeSecretKey

	pkg.Memory = redis.NewClient(&redis.Options{
		Addr:     pkg.RedisAddr,
		Password: pkg.RedisPassword,
	})
	if err := pkg.Memory.Ping(ctx).Err(); err != nil {
		log.Fatalln(err)
	}

	pkg.Mail = sendgrid.NewSendClient(pkg.SendGridSecretKey)

	// pkg.Config = pkg.ConfigData{
	// 	Memory: redisClient,
	// 	Mail:   sendGridClient,
	// }

	r := gin.Default()
	r.Static("/static", "./web/public")

	v1.Set(r.Group("/api/v1"))

	r.Run(":8080")
}
