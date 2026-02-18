package rolesmodel

import (
	"fmt"
	"wms/pkg/utils"
)

type Role struct {
	ID         		uint64    `gorm:"primaryKey;" json:"id"`
	Name      		string    `gorm:"uniqueIndex;size:40;not null" json:"name"`
}

func (Role) TableName() string {
	return "roles"
}

func (Role) GetUpdateMap(setRequest any) (map[string]any, error) {
	var err error
	fields :=  map[string]any{}

	fields["name"], err = utils.GetStructFieldByString(setRequest, "Name")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	return fields, nil
}

func New() *Role {return &Role{}}