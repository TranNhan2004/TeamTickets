package permissions

import "time"

type Permission struct {
	ID          string     `gorm:"primaryKey;"`
	Code        string     `gorm:"column:code;"`
	Description *string    `gorm:"column:description;"`
	IsDeleted   bool       `gorm:"column:is_deleted;"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;"`
	CreatedAt   time.Time  `gorm:"column:created_at;"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;"`
}
