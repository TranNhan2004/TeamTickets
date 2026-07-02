package health

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type DatabaseChecker struct {
	db *gorm.DB
}

func NewDatabaseChecker(db *gorm.DB) *DatabaseChecker {
	return &DatabaseChecker{
		db: db,
	}
}

func (c *DatabaseChecker) Name() string {
	return "database"
}

func (c *DatabaseChecker) Check(ctx context.Context) error {
	var result int

	err := c.db.WithContext(ctx).
		Raw("SELECT 1").
		Scan(&result).
		Error

	if err != nil {
		return fmt.Errorf("database check failed: %w", err)
	}

	if result != 1 {
		return fmt.Errorf("database check returned invalid result")
	}

	return nil
}
