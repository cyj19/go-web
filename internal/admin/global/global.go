package global

import (
	"github.com/cyj19/go-web/internal/pkg/config"
	"github.com/cyj19/go-web/internal/pkg/logger"

	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

/*
	ćšć±ćé
*/

var (
	Box      *config.CustomConfBox
	Conf     *config.Configuration
	Log      *logger.GormZapLogger
	DB       *gorm.DB
	RedisIns *redis.Client
	Enforcer *casbin.Enforcer
)
