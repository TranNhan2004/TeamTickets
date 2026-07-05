package shared

import (
	"time"

	"github.com/google/uuid"
)

type AuditModel struct {
	ActorID   *uuid.UUID
	AuditTime time.Time
}
