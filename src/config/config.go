package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//GetDB connection
func GetDB() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "T3rr1613"
	dbName := "netcoreapidb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	return db, err
}

//SetDB for setting the db configuration
// func SetDB() {
// 	db, err := GetDB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	DbH := models.DbHandler{
// 		Db: db,
// 	}
// }
