package orm

import (
	"context"
	"errors"
	"github.com/aknEvrnky/notification-system/internal/application/core/domain"
	"gorm.io/gorm"
	"time"
)

type user struct {
	gorm.Model
	ID          string `gorm:"primarykey"`
	name        string
	email       string
	phone       string
	deviceToken string
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindById(ctx context.Context, userId string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	usr := user{ID: userId}

	u.db.First(&usr)

	if u.db.Error != nil {
		if errors.Is(u.db.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, nil // User not found
		}
		return domain.User{}, u.db.Error // Other database error
	}

	return domain.User{
		Id:          usr.ID,
		Name:        usr.name,
		Email:       usr.email,
		Phone:       usr.phone,
		DeviceToken: usr.deviceToken,
	}, nil
}
