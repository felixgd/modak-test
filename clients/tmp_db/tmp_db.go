package tmpdb

import (
	"fmt"
	"modak-test/domain/user"
)

type DB struct{}

var users_db []user.User

func (db *DB) PopulateUsers(numberOfUsers int) {
	for i := 0; i < numberOfUsers; i++ {
		newUser := user.User{
			ID:            i + 1,
			Email:         fmt.Sprintf("test%v@test.com", (i + 1)),
			Notifications: make(map[string]user.NotifiedUser),
		}
		users_db = append(users_db, newUser)
	}
}

func (db *DB) FindUserByID(userID int) *user.User {
	for _, user := range users_db {
		if user.ID == userID {
			return &user
		}
	}

	return nil
}

func (db *DB) GetAllUsers() []user.User {
	return users_db
}
