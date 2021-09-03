package global

import (
	"github.com/vagaryer/go-web/internal/pkg/config"
	"github.com/vagaryer/go-web/internal/pkg/logger"

	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

/*
	全局变量
*/

var (
	Box      *config.CustomConfBox
	Conf     *config.Configuration
	Log      *logger.GormZapLogger
	DB       *gorm.DB
	RedisIns *redis.Client
	Enforcer *casbin.Enforcer
)
