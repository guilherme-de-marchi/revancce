package pkg

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	ServerDomain = os.Getenv("SERVER_DOMAIN")

	ProductPurchaseRedirectSuccess = os.Getenv("PRODUCT_PURCHASE_REDIRECT_SUCCESS")
	ProductPurchaseRedirectCancel  = os.Getenv("PRODUCT_PURCHASE_REDIRECT_CANCEL")

	StripeSecretKey     = os.Getenv("STRIPE_SECRET_KEY")
	StripeWebhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")

	SendGridSecretKey      = os.Getenv("SENDGRID_SECRET_KEY")
	SendGridPurchasesEmail = os.Getenv("SENDGRID_PURCHASES_EMAIL")
	SendGridPurchasesName  = os.Getenv("SENDGRID_PURCHASES_NAME")

	RedisAddr     = os.Getenv("REDIS_ADDR")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
)
