package repository 

import "database/sql"

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb{
	retyrn &TransactionRepositoryDb{db: db}
}