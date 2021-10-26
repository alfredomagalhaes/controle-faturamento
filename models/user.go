package models

import (
	"errors"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrEmailNotFound = errors.New("email address not found")
var ErrDBConn = errors.New("connection error. Please retry")
var ErrInvalidCredentials = errors.New("invalid login credentials. Please try again")

//Token JWT claims struct
type Token struct {
	UserID uuid.UUID
	jwt.StandardClaims
}

//User a struct to rep user User
type User struct {
	Base
	Email    string `json:"email";gorm:"unique"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token";gorm:"-"`
}

//Login method to valid user, password and token
func (u *User) Login() error {

	user := &User{}
	err := config.MI.DB.Table("users").Where("email = ?", u.Email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrEmailNotFound
		}
		return ErrDBConn
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return ErrInvalidCredentials
	}
	//Worked! Logged In
	user.Password = ""

	return nil
}
