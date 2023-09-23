package pkg

import (
	"github.com/redis/go-redis/v9"
	"github.com/sendgrid/sendgrid-go"
)

var (
	Memory *redis.Client
	Mail   *sendgrid.Client
)
