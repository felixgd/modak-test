package gateway

import (
	"fmt"
	"modak-test/domain/notification"
	userDomain "modak-test/domain/user"
)

// Gateway represents the service that sends notifications.
type Gateway struct{}

// Send sends a notification to a user.
func (g *Gateway) Send(user userDomain.User, notification notification.Notification) {
	fmt.Printf("Sending %s notification to userID %v and email %s: %s\n", notification.Type, user.ID, user.Email, notification.Message)
}
