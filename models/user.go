package models

import (
	"easy-search/db"
)

type User struct {
	ID    uint   `gorm:"primarykey"`
	Qq    string `gorm:"unique"`
	Phone string `gorm:"unique"`
}

func GetUserPhone(user User) (*User, error) {
	result := db.Conn.Model(&User{}).Where("phone = ?", user.Phone).Find(&user)
	// api层做过拦截，所以此处user.Phone必定有值
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserQq(user User) (*User, error) {
	result := db.Conn.Model(&User{}).Where("qq = ?", user.Qq).Find(&user)
	// api层做过拦截，所以此处user.Phone必定有值
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}
