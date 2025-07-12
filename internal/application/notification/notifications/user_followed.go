package notifications

import (
	"fmt"
	"github.com/aknEvrnky/notification-system/internal/application/core/domain"
	"github.com/aknEvrnky/notification-system/internal/application/notification"
	"github.com/aknEvrnky/notification-system/internal/application/notification/messages"
)

type UserFollowedNotification struct {
	User         domain.User
	FollowerUser domain.User
}

func NewUserFollowedNotification(user, followerUser domain.User) *UserFollowedNotification {
	return &UserFollowedNotification{
		User:         user,
		FollowerUser: followerUser,
	}
}

func (n *UserFollowedNotification) ToMail() (*messages.MailMessage, error) {
	body := fmt.Sprintf("Hello, %s started following you,", n.FollowerUser.Name)

	return &messages.MailMessage{
		To:      n.User.Email,
		Subject: "New Follower Notification",
		Body:    body,
	}, nil
}

func (n *UserFollowedNotification) Channels() []notification.NotificationChannel {
	return []notification.NotificationChannel{
		notification.ChannelMail,
	}
}
