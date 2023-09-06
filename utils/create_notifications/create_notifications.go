package createnotifications

import (
	"math/rand"
	"modak-test/domain/notification"
	"modak-test/domain/user"
	ratelimiterrules "modak-test/utils/rate_limiter_rules"
	"time"
)

// CreateRandomNotification is used for simulating incoming notifications traffic
func CreateRandomNotification(ch chan notification.Notification, users []user.User, rules map[string]ratelimiterrules.RateLimitRule) {
	notificationTypes := make([]string, 0, len(rules))
	for k := range rules {
		notificationTypes = append(notificationTypes, k)
	}

	for true {
		time.Sleep(time.Second)
		randomID := rand.Intn(len(users)-1) + 1
		randomType := rand.Intn(len(notificationTypes) - 1)
		newNotification := notification.Notification{
			Type:    notificationTypes[randomType],
			UserID:  randomID,
			Message: "This is a test",
		}

		ch <- newNotification
	}
}
