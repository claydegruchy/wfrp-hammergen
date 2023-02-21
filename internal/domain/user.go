package domain

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const (
	UserNotFoundError = iota
	UserAlreadyExistsError
	UserInternalError
	UserIncorrectPasswordError
	UserInvalidArgumentsError
	UserSendEmailError
	UserUnauthorizedError
)

type User struct {
	Id                 string
	Username           string
	Admin              bool
	SharedAccountNames []string
	SharedAccountIds   []string
	Password           string
	PasswordHash       []byte
	CreatedOn          time.Time
	LastAuthOn         time.Time
}

func (u *User) Copy() *User {
	if u == nil {
		return nil
	}
	uCopy := User{}
	uCopy.Id = strings.Clone(u.Id)
	uCopy.Username = strings.Clone(u.Username)
	uCopy.Admin = u.Admin

	if u.SharedAccountNames != nil {
		uCopy.SharedAccountNames = make([]string, len(u.SharedAccountNames))
		copy(uCopy.SharedAccountNames, u.SharedAccountNames)
	} else {
		uCopy.SharedAccountNames = nil
	}

	if u.SharedAccountIds != nil {
		uCopy.SharedAccountIds = make([]string, len(u.SharedAccountIds))
		copy(uCopy.SharedAccountIds, u.SharedAccountIds)
	} else {
		uCopy.SharedAccountIds = nil
	}

	uCopy.Password = strings.Clone(u.Password)

	if u.PasswordHash != nil {
		uCopy.PasswordHash = make([]byte, len(u.PasswordHash))
		copy(uCopy.PasswordHash, u.PasswordHash)
	}

	uCopy.LastAuthOn = u.LastAuthOn.UTC()
	uCopy.CreatedOn = u.CreatedOn.UTC()

	return &uCopy
}

func EmptyUser() *User {
	u := &User{}
	u.SharedAccountNames = make([]string, 0)
	u.SharedAccountIds = make([]string, 0)
	u.PasswordHash = make([]byte, 0)

	return u
}

type UserService interface {
	Get(ctx context.Context, c *Claims, id string) (*User, *UserError)
	Exists(ctx context.Context, username string) (bool, *UserError)
	Create(ctx context.Context, u *User) (*User, *UserError)
	Update(ctx context.Context, c *Claims, u *User) (*User, *UserError)
	UpdateCredentials(ctx context.Context, c *Claims, currentPasswd string, u *User) (*User, *UserError)
	UpdateClaims(ctx context.Context, c *Claims, u *User) (*User, *UserError)
	Delete(ctx context.Context, c *Claims, id string) *UserError
	List(ctx context.Context, c *Claims) ([]*User, *UserError)
	Authenticate(ctx context.Context, username string, password string) (u *User, ue *UserError)
	SendResetPassword(ctx context.Context, username string) *UserError
	ResetPassword(ctx context.Context, token string, newPassword string) *UserError
}

type UserError struct {
	Type int
	Err  error
}

func (e *UserError) Unwrap() error {
	return e.Err
}

func (e *UserError) Error() string {
	return fmt.Sprintf("user error, %s", e.Err)
}
