package db

import (
	"crypto/tls"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func InitMysql() *bun.DB {
	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: GetHost(),
	})

	sqldb, err := sql.Open("mysql", GetDsn())
	if err != nil {
		panic(err)
	}
	sqldb.SetMaxIdleConns(1)
	sqldb.SetMaxOpenConns(1)
	db := bun.NewDB(sqldb, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return db
}
