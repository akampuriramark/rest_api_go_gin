package config

import (
	"database/sql"
	"gorestapimysql/src/models"
)

//GetDB connection
func GetDB() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "T3rr1bl3"
	dbName := "netcoreapidb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}

//SetDB for setting the db configuration
func SetDB() {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	DbH := models.DbHandler{
		Db: db,
	}
}
