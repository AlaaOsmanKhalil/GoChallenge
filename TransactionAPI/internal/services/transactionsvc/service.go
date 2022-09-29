package transactionsvc

import (
	"TransactionAPI/internal/repositories/transaction"
	"context"
)

type IService interface {
	Create(ctx context.Context, payload CreatePld) (*transaction.Model, error)
	GetAll(ctx context.Context) (*[]transaction.Model, error)
}
