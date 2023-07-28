package db

import (
	"os"
)

func GetDsn() string {
	return os.Getenv("tidb_dsn")
}

func GetHost() string {
	return os.Getenv("tidb_host")
}
