package gateway

import (
	"fmt"
	userDomain "modak-test/domain/user"
)

// Gateway represents the service that sends notifications.
type Gateway struct{}

// Send sends a notification to a user.
func (g *Gateway) Send(user userDomain.User, message string) {
	fmt.Printf("Sending message to user %v: %s\n", user.ID, message)
}
