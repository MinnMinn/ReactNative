package driver

import (
	"database/sql"
	"rest-api-mysql/config"
)

func DbConn() (db *sql.DB) {
	db, err := sql.Open(config.DbDriver, config.DbUser+":"+config.DbPass+"@/"+config.DbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
