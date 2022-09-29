package transactionsvc

import (
	"TransactionAPI/internal/repositories/transaction"
	"context"
	"go.uber.org/zap"
)

type DefaultService struct {
	log             *zap.SugaredLogger
	transactionRepo transaction.IRepository
}

func NewDefaultService(log *zap.SugaredLogger, tr transaction.IRepository) *DefaultService {
	return &DefaultService{
		log:             log,
		transactionRepo: tr,
	}
}

func (s *DefaultService) Create(ctx context.Context, payload CreatePld) (*transaction.Model, error) {
	dbModel := toRepoCreateModel(payload)
	result, err := s.transactionRepo.Create(ctx, &dbModel)

	if err != nil {
		s.log.Errorf("Fail to create a new transaction")
		return nil, err
	}

	return result, nil
}

func (s *DefaultService) GetAll(ctx context.Context) (*[]transaction.Model, error) {
	result, err := s.transactionRepo.GetAll(ctx)

	if err != nil {
		s.log.Errorf("Cannot get all transactions")
		return nil, err
	}

	if err != nil {
		s.log.Errorf("Failed to map transactions model to getres model")
		return nil, err
	}
	return result, nil
}
