package pkg

import (
	"crypto/rsa"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/sendgrid/sendgrid-go"
)

var (
	Log      *log.Logger
	Database *pgxpool.Pool
	Memory   *redis.Client
	Mail     *sendgrid.Client
	RSA      *rsa.PrivateKey
)
