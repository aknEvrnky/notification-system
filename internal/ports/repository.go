package ports

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/core/domain"
)

type UserRepository interface {
	FindById(ctx context.Context, userId string) (domain.User, error)
}
