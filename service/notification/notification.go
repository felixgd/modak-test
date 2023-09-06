package notification

import (
	"fmt"
	"time"

	"modak-test/clients/gateway"
	tmpdb "modak-test/clients/tmp_db"
	"modak-test/domain/notification"
	userDomain "modak-test/domain/user"
	ratelimiterrules "modak-test/utils/rate_limiter_rules"
)

type NotificationService interface {
	Send(string, string, string)
}

// NotificationServiceImpl implements the NotificationService interface.
type NotificationServiceImpl struct {
	gateway        gateway.Gateway
	db             tmpdb.DB
	rateLimitRules map[string]ratelimiterrules.RateLimitRule
}

// NewNotificationService creates a new NotificationServiceImpl instance.
func NewNotificationService(gateway gateway.Gateway, db tmpdb.DB, rules map[string]ratelimiterrules.RateLimitRule) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		gateway:        gateway,
		db:             db,
		rateLimitRules: rules,
	}
}

// Send sends a notification if it does not exceed the rate limit.
func (s *NotificationServiceImpl) Send(notification notification.Notification) {
	rule, ok := s.rateLimitRules[notification.Type]
	if !ok {
		fmt.Printf("Unknown notification type: %s\n", notification.Type)
		return
	}

	user := s.db.FindUserByID(notification.UserID)

	if user == nil {
		fmt.Printf("Unknown user: %v\n", user.ID)
		return
	}

	if s.isRateLimited(user, notification.Type, rule.MaxCount, rule.Duration) {
		fmt.Printf("Rate limited: %s\n", notification.Message)
		return
	}

	s.gateway.Send(*user, notification.Message)
	s.incrementCounter(user, notification.Type)
}

// isRateLimited checks if the user has exceeded the rate limit for a given notification type.
func (s *NotificationServiceImpl) isRateLimited(user *userDomain.User, notificationType string, maxCount int, duration time.Duration) bool {
	now := time.Now()

	notification, ok := user.Notifications[notificationType]
	if !ok {
		user.Notifications[notificationType] = userDomain.NotifiedUser{
			Counter:       user.Notifications[notificationType].Counter,
			LastTimestamp: now,
		}
		return false
	}

	test := now.Sub(notification.LastTimestamp)

	if test <= duration {
		return user.Notifications[notificationType].Counter >= maxCount
	}

	return false
}

// incrementCounter increments the counter for a given user and notification type.
func (s *NotificationServiceImpl) incrementCounter(user *userDomain.User, notificationType string) {
	now := time.Now()
	user.Notifications[notificationType] = userDomain.NotifiedUser{
		Counter:       user.Notifications[notificationType].Counter + 1,
		LastTimestamp: now,
	}
}
