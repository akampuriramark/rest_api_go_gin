package models

import (
	"database/sql"
	"gorestapimysql/src/entities"
)

//DbHandler to call when executing stored procedures.
type DbHandler struct {
	Db *sql.DB
}

// ExecuteDataSet to call when executing stored procedures expecting results.
func (db DbHandler) ExecuteDataSet(storedProcedure string) (*sql.Rows, error) {
	rows, err := db.Db.Query(storedProcedure)
	return rows, err
}

// ExecuteNonQuery to call when executing stored procedures expecting no results.
func (db DbHandler) ExecuteNonQuery(storedProcedure string, transaction entities.Transaction) (sql.Result, error) {
	result, err := db.Db.Exec(storedProcedure,
		transaction.CustomerRef, transaction.AgentCode, transaction.AgentID, transaction.Amount, transaction.PaymentDate,
		transaction.CompletedAtVendor, transaction.SentToUtility, transaction.UtilitySentDate, transaction.UtilityReference, transaction.RecordDate)
	return result, err
}
