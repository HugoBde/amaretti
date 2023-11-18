package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	. "hugobde.dev/amaretti/password"
)

type User struct {
	UUID     uuid.UUID
	username string
	password Password
}

type Session struct {
	uuid       uuid.UUID
	user       *User
	expiryTime time.Time
}

var UserDB map[string]*User = make(map[string]*User)

var SessionDB map[uuid.UUID]Session = make(map[uuid.UUID]Session)

func NewUser(username string, password string) error {
	// Ensure password is ok
	if !IsPasswordValid(password) {
		return errors.New("Invalid Password")
	}

	var newUser User
	var err error

	// Create user password hash
	if newUser.password, err = CreatePassword(password); err != nil {
		return err
	}

	// At this point we are all good, save the username
	newUser.username = username
	newUser.UUID = uuid.New()

	// Store user in our mock db
	UserDB[username] = &newUser

	return nil
}

func createSession(user *User) Session {
	newSession := Session{
		uuid:       uuid.New(),
		user:       user,
		expiryTime: time.Now().Add(time.Minute),
	}

	time.AfterFunc(time.Minute, func() {
		delete(SessionDB, newSession.uuid)
	})

	return newSession
}

func AuthenticateUser(username string, password string) bool {
	user, ok := UserDB[username]

	if !ok {
		return false
	}

	if !VerifyPassword(password, user.password) {
		return false
	}

	createSession(user)

	return VerifyPassword(password, user.password)
}
