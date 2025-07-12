package api

import (
	"context"
	"github.com/aknEvrnky/notification-system/internal/application/core/domain"
)

func (a *Application) CreateUser(ctx context.Context, user domain.User) error {
	return a.UserRepository.Create(ctx, user)
}

func (a *Application) UpdateUser(ctx context.Context, user domain.User) error {
	return a.UserRepository.Update(ctx, user)
}

func (a *Application) DeleteUser(ctx context.Context, id string) error {
	return a.UserRepository.Delete(ctx, id)
}
