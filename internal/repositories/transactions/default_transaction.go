package transactions

import (
	Db "Challenge/internal/adapters/db"

	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

func (s *DefaultTransactionService) GetAllTransactions() ([]Transaction, error) {
	db := Db.CreateDatabase()
	var trans = []Transaction{}
	db.NewSelect().Model(&trans).Scan(context.Background())
	if trans != nil {

		return trans, nil
	} else {

		return nil, errors.New("no transactions found")
	}
}

func (s *DefaultTransactionService) CreateTransaction(trans *Transaction) (Transaction, error) {
	db := Db.CreateDatabase()
	trans.Id = uuid.New()
	trans.CreatedAt = time.Now().String()
	trans.Status = false
	if _, err := db.NewInsert().Model(trans).Exec(context.Background()); err != nil {
		return *trans, err
	}
	Produce(trans)
	return *trans, nil

}

func (s *DefaultTransactionService) UpdateTransaction(JsonTransaction string) (bool, error) {
	db := Db.CreateDatabase()
	var transaction Transaction
	json.Unmarshal([]byte(JsonTransaction), &transaction)
	transaction.Status = true
	if _, err := db.NewUpdate().Model(&transaction).WherePK().Exec(context.Background()); err != nil {
		return false, nil
	}
	return true, nil
}
