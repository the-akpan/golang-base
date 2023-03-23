package models

import (
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `json:"email" gorm:"size:100;index:,unique;not null;"`
	Mobile   string `json:"mobile" gorm:"size:15;index:,unique;not null;"`
	Username string `json:"username" gorm:"size:15;index:,unique;not null;"`
	Password string `json:"password" gorm:"size:30;:,"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {

	if err := user.Base.BeforeCreate(tx); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return err
	}

	if len(user.Password) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("Password", pw)

	return nil
}
