package models

import "gorestapimysql/src/entities"

//GetTransactions to query for all transactions
func (db DbHandler) GetTransactions() ([]entities.Transaction, error) {
	rows, err := db.ExecuteDataSet("getTransactions", []string{})
	if err != nil {
		return nil, err
	}
	transactions := []entities.Transaction{}
	for rows.Next() {
		transaction := entities.Transaction{}
		errar := rows.Scan(transaction.Idnumber, transaction.Utility, transaction.Vendorcode, transaction.Amount)
		if errar != nil {
			return nil, errar
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
