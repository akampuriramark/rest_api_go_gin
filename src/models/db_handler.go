package models

import "database/sql"

//DbHandler to call when executing stored procedures.
type DbHandler struct {
	Db *sql.DB
}

// ExecuteDataSet to call when executing stored procedures expecting results.
func (db DbHandler) ExecuteDataSet(storedProcedure string, parameters []string) (*sql.Rows, error) {
	rows, err := db.Db.Query(storedProcedure, parameters)
	return rows, err
}
