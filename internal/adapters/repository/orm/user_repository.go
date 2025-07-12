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
	Name        string
	Email       string
	Phone       string
	DeviceToken string
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
		Name:        usr.Name,
		Email:       usr.Email,
		Phone:       usr.Phone,
		DeviceToken: usr.DeviceToken,
	}, nil
}

func (u *UserRepository) Create(ctx context.Context, usr domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	record := user{
		ID:          usr.Id,
		Name:        usr.Name,
		Email:       usr.Email,
		Phone:       usr.Phone,
		DeviceToken: usr.DeviceToken,
	}

	return u.db.WithContext(ctx).Create(&record).Error
}

func (u *UserRepository) Update(ctx context.Context, usr domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	record := user{
		ID:          usr.Id,
		Name:        usr.Name,
		Email:       usr.Email,
		Phone:       usr.Phone,
		DeviceToken: usr.DeviceToken,
	}

	return u.db.WithContext(ctx).Model(&user{}).Where("id = ?", usr.Id).Updates(record).Error
}

func (u *UserRepository) Delete(ctx context.Context, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return u.db.WithContext(ctx).Where("id = ?", userId).Delete(&user{}).Error
}
