package lib

import (
	"github.com/jmoiron/sqlx"
	"github.com/miajio/www/bin"
	"go.uber.org/zap"
)

var (
	Log       *zap.SugaredLogger
	DBCfg     bin.MysqlCfgParam
	ServerCfg HttpServerParam
	DB        *sqlx.DB
)
