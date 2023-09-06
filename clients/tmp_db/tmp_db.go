package tmpdb

import "modak-test/domain/user"

type DB struct{}

var USERS_DB []user.User

func (db *DB) PopulateUsers(numberOfUsers int) {
	for i := 0; i < numberOfUsers; i++ {
		newUser := user.User{
			ID:            i + 1,
			Notifications: make(map[string]user.NotifiedUser),
		}
		USERS_DB = append(USERS_DB, newUser)
	}
}

func (db *DB) FindUserByID(userID int) *user.User {
	for _, user := range USERS_DB {
		if user.ID == userID {
			return &user
		}
	}

	return nil
}
