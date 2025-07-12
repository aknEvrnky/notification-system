package dispatcher

import (
	"context"
	"errors"
	"github.com/aknEvrnky/notification-system/internal/application/notification"
	"github.com/aknEvrnky/notification-system/internal/application/notification/notifications"
	"github.com/aknEvrnky/notification-system/internal/ports"
)

type UserFollowedHandler struct {
	userRepository ports.UserRepository
}

func NewUserFollowedHandler(userRepo ports.UserRepository) *UserFollowedHandler {
	return &UserFollowedHandler{userRepository: userRepo}
}

func (h *UserFollowedHandler) Handle(ctx context.Context, payload map[string]any) ([]notification.Notification, error) {
	// 1. Gelen payload'dan userId ve followerId çek
	userID, ok1 := payload["user_id"].(string)
	followerID, ok2 := payload["follower_id"].(string)
	if !ok1 || !ok2 {
		return nil, errors.New("invalid payload: user_id and follower_id are required")
	}

	// 2. Domain bilgilerini al
	user, err := h.userRepository.FindById(ctx, userID)
	if err != nil {
		return nil, err
	}

	follower, err := h.userRepository.FindById(ctx, followerID)
	if err != nil {
		return nil, err
	}

	// 3. Notification oluştur
	notif := notifications.NewUserFollowedNotification(user, follower)

	return []notification.Notification{notif}, nil
}
