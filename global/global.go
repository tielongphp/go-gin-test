package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"go-gin-test/context"
)

var (
	DB         *gorm.DB
	REDIS      *redis.Client
	CTX_CONFIG *context.Config
)
