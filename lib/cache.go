package lib

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	Log         *zap.SugaredLogger
	DBCfg       MysqlCfgParam
	RedisCfg    = &RedisCfgParam{}
	RedisClient *redis.Client
	ServerCfg   HttpServerParam
	EmailCfg    = &EmailCfgParam{}
	DB          *sqlx.DB
)
