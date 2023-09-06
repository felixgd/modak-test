package main

import (
	"modak-test/clients/gateway"
	tmpdb "modak-test/clients/tmp_db"
	"modak-test/domain/notification"
	notificationService "modak-test/service/notification"
	createnotifications "modak-test/utils/create_notifications"
	ratelimiterrules "modak-test/utils/rate_limiter_rules"
	"time"
)

func main() {
	gateway := gateway.Gateway{}
	db := tmpdb.DB{}
	notificationChannel := make(chan notification.Notification)

	rateLimitRules := map[string]ratelimiterrules.RateLimitRule{
		"status":    {MaxCount: 2, Duration: time.Minute},
		"news":      {MaxCount: 1, Duration: 24 * time.Hour},
		"marketing": {MaxCount: 3, Duration: time.Hour},
	}

	db.PopulateUsers(3)

	go createnotifications.CreateRandomNotification(notificationChannel, db.GetAllUsers(), rateLimitRules)

	service := notificationService.NewNotificationService(gateway, db, rateLimitRules)

	for notificationToSend := range notificationChannel {
		service.Send(notificationToSend)
	}

	close(notificationChannel)
}
