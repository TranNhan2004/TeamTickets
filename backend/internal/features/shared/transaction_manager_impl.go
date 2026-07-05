package shared

import (
	"context"

	"gorm.io/gorm"
)

type txKey struct{}

type transactionManagerImpl struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) TransactionManager {
	return &transactionManagerImpl{db: db}
}

func (t *transactionManagerImpl) WithinTransaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, txKey{}, tx)
		return fn(txCtx)
	})
}

func DBFromContext(ctx context.Context, fallback *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}

	return fallback
}
