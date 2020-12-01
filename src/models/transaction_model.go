package models

import (
	"gorestapimysql/src/entities"
)

//GetTransactions to query for all transactions
func (db DbHandler) GetTransactions() ([]entities.Transaction, error) {
	rows, err := db.ExecuteDataSet("call getTransactions()")
	if err != nil {
		return nil, err
	}
	transactions := []entities.Transaction{}
	for rows.Next() {
		transaction := entities.Transaction{}
		errar := rows.Scan(&transaction.CustomerRef, &transaction.AgentCode, &transaction.AgentID, &transaction.Amount, &transaction.PaymentDate, &transaction.CompletedAtVendor,
			&transaction.SentToUtility, &transaction.UtilitySentDate, &transaction.UtilityReference, &transaction.RecordDate)
		if errar != nil {
			return nil, errar
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

//SaveTransaction saves transaction to db
func (db DbHandler) SaveTransaction(transaction entities.Transaction) (int64, error) {
	result, err := db.ExecuteNonQuery("call saveTransaction(?,?,?,?,?,?,?,?,?,?)", transaction)
	if err != nil {
		return 0, err
	}
	lastID, erra := result.LastInsertId()
	if erra != nil {
		return 0, err
	}
	return lastID, nil

}
