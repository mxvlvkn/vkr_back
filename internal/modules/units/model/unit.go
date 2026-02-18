package unitsmodel

import (
	"fmt"
	"wms/pkg/utils"
)

type Unit struct {
	ID         		uint64    `gorm:"primaryKey;" json:"id"`
	Name      		string    `gorm:"size:40;not null" json:"name"`
	Code   			uint   	  `gorm:"uniqueIndex;not null" json:"code"`
	Sign       		string    `gorm:"size:10;not null" json:"sign"`
}

func (Unit) TableName() string {
	return "units"
}

func (Unit) GetUpdateMap(setRequest any) (map[string]any, error) {
	var err error
	fields :=  map[string]any{}

	fields["name"], err = utils.GetStructFieldByString(setRequest, "Name")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["code"], err = utils.GetStructFieldByString(setRequest, "Code")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["sign"], err = utils.GetStructFieldByString(setRequest, "Sign")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	return fields, nil
}

func New() *Unit {return &Unit{}}