package usermodel

import (
	"errors"
	"food-delivery/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email"`
	Password        string        `json:"-" gorm:"column:password"`
	Salt            string        `json:"-" gorm:"column:salt"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	FirstName       string        `json:"first_name" gorm:"column:first_name"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"role" gorm:"column:role"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string { return "users" }

func (u *User) getUserId() int {
	return u.ID
}
func (u *User) getEmail() string {
	return u.Email
}
func (u *User) getRole() string {
	return u.Role
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email"`
	Password        string        `json:"password" gorm:"column:password"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	FirstName       string        `json:"first_name" gorm:"column:first_name"`
	Role            string        `json:"_" gorm:"column:role"`
	Salt            string        `json:"-" gorm:"column:salt"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string { return User{}.TableName() }

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string { return User{}.TableName() }

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already exists"),
		"email has already exists",
		"ErrEmailExisted",
	)
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
)
