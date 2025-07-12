package ports

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/core/domain"
)

type UserRepository interface {
	FindById(ctx context.Context, userId string) (domain.User, error)
	Create(ctx context.Context, user domain.User) error
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, userId string) error
}
