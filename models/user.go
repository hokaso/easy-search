package models

import (
	"easy-search/db"
)

type SenQq struct {
	ID    uint `gorm:"primarykey"`
	Qq    string
	Phone string
}

func GetUserPhone(user SenQq) (*SenQq, error) {
	result := db.Conn.Model(&SenQq{}).Where("phone = ?", user.Phone).Find(&user)
	// api层做过拦截，所以此处user.Phone必定有值
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserQq(user SenQq) (*SenQq, error) {
	result := db.Conn.Model(&SenQq{}).Where("qq = ?", user.Qq).Find(&user)
	// api层做过拦截，所以此处user.Phone必定有值
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}
