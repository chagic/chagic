package model

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base

	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) GenerateToken() (string, error) {
	jwtTTL, err1 := strconv.Atoi(os.Getenv("JWT_TTL"))
	if err1 != nil {
		panic("Generate token error")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = u
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(jwtTTL)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return s, err
}
