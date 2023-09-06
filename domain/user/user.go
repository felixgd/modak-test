package user

import "time"

type User struct {
	ID            int
	Notifications map[string]NotifiedUser
}

type NotifiedUser struct {
	Counter       int
	LastTimestamp time.Time
}
