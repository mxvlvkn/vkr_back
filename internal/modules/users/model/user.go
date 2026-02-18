package usersmodel

import (
	"fmt"
	"wms/pkg/utils"
)

type User struct {
	ID         		uint64    `gorm:"primaryKey;" json:"id"`
	Login      		string    `gorm:"uniqueIndex;size:40;not null" json:"login"`
	PasswordHash   	string    `gorm:"column:password_hash;size:40;not null" json:"-"`
	Name       		string    `gorm:"size:40;not null" json:"name"`
	Surname    		string    `gorm:"size:40;not null" json:"surname"`
	Patronymic 		string    `gorm:"size:40;not null" json:"patronymic"`
	RoleID       		uint64    `gorm:"not null" json:"role_id"`
}

func (User) TableName() string {
	return "users"
}

func (User) GetUpdateMap(setRequest any) (map[string]any, error) {
	fields, err := utils.FillMapFromStruct(
		setRequest,
		&map[string]string{
			"login": 		"Login",
			"name": 		"Name",
			"surname": 		"Surname",
			"patronymic": 	"Patronymic",
			"role_id": 		"RoleID",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	return fields, nil
}

func New() *User {return &User{}}