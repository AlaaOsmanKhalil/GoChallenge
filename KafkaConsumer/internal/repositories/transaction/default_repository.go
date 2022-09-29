package transaction

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"time"
)

type DatabaseRepository struct {
	log *zap.SugaredLogger
	db  *bun.DB
}

func NewDatabaseRepository(log *zap.SugaredLogger, conn *bun.DB) *DatabaseRepository {
	return &DatabaseRepository{
		log: log,
		db:  conn,
	}
}

func (r *DatabaseRepository) Create(ctx context.Context, model *Model) (*Model, error) {
	model.ID = uuid.New()
	model.CreatedAt = time.Now()
	_, err := r.db.NewInsert().Model(model).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *DatabaseRepository) GetAll(ctx context.Context) (*[]Model, error) {
	var transactions = &[]Model{}
	err := r.db.NewSelect().Model(transactions).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
