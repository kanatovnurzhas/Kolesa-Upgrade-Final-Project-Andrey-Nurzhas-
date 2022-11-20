package models

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Name       string `json:"name"`
	TelegramId int64  `json:"telegram_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ChatId     int64  `json:"chat_id"`
}

type UserModel struct {
	Db *gorm.DB
}

func (u *UserModel) Create(user User) error {

	result := u.Db.Create(&user)

	return result.Error
}

func (u *UserModel) FindOne(telegramId int64) (*User, error) {
	existUser := User{}

	result := u.Db.First(&existUser, User{TelegramId: telegramId})
	if result.Error != nil {
		return nil, result.Error
	}

	return &existUser, nil
}

func (u *UserModel) FindAllUsers() []User {
	var existUser []User
	result := u.Db.Find(&existUser)
	if result.Error != nil {
		log.Printf("Ошибка при получении пользователей %v", result.Error)
	}
	fmt.Println(existUser[0].TelegramId)
	return existUser

}
