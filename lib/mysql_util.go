package lib

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mysqlUtilImpl struct{}

type mysqlUtil interface {
	DSN(param MysqlCfgParam) string
	Connect(param MysqlCfgParam) (*sqlx.DB, error)
}

// Mysql 配置参数
type MysqlCfgParam struct {
	Host      string `toml:"host"`      // 数据库地址 127.0.0.1:3306
	User      string `toml:"user"`      // 数据库用户名
	Password  string `toml:"password"`  // 数据库密码
	Database  string `toml:"database"`  // 数据库名
	Charset   string `toml:"charset"`   // 数据库字符集 utf8mb4
	ParseTime string `toml:"parseTime"` // 是否分析时间 True
	Loc       string `toml:"loc"`       // loc Local
}

var MysqlUtil mysqlUtil = (*mysqlUtilImpl)(nil)

func (*mysqlUtilImpl) DSN(param MysqlCfgParam) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", param.User, param.Password, param.Host, param.Database)
	params := make(map[string]string, 0)
	if param.Charset != "" {
		params["charset"] = param.Charset
	}
	if param.ParseTime != "" {
		params["parseTime"] = param.ParseTime
	}
	if param.Loc != "" {
		params["loc"] = param.Loc
	}

	vals := make([]string, 0)
	for k, v := range params {
		vals = append(vals, k+"="+v)
	}
	ps := strings.Join(vals, "&")
	if ps != "" {
		dsn = dsn + "?" + ps
	}
	return dsn
}

func (mu *mysqlUtilImpl) Connect(param MysqlCfgParam) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", mu.DSN(param))
	if err != nil {
		return nil, err
	}
	return db, nil
}
