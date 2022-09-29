package transactionsvc

import (
	"TransactionAPI/internal/repositories/transaction"
	"github.com/google/uuid"
	"time"
)

func toRepoCreateModel(data CreatePld) transaction.Model {
	transactionId := uuid.NewString()
	now := time.Now().UTC()
	return transaction.Model{
		ID:        transactionId,
		Amount:    data.Amount,
		Currency:  data.Currency,
		CreatedAt: now,
	}
}
