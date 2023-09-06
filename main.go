package main

import (
	"modak-test/clients/gateway"
	tmpdb "modak-test/clients/tmp_db"
	"modak-test/domain/notification"
	notificationService "modak-test/service/notification"
	ratelimiterrules "modak-test/utils/rate_limiter_rules"
	"time"
)

func main() {
	gateway := gateway.Gateway{}
	rateLimitRules := map[string]ratelimiterrules.RateLimitRule{
		"status":    {MaxCount: 2, Duration: time.Minute},
		"news":      {MaxCount: 1, Duration: time.Hour},
		"marketing": {MaxCount: 3, Duration: time.Hour},
	}

	db := tmpdb.DB{}
	db.PopulateUsers(3)

	service := notificationService.NewNotificationService(gateway, db, rateLimitRules)

	notificationToSend := notification.Notification{
		Type:    "news",
		UserID:  1,
		Message: "test",
	}

	service.Send(notificationToSend)
	service.Send(notificationToSend)
	service.Send(notificationToSend)
	service.Send(notificationToSend)
	service.Send(notificationToSend)
}
