package manufacturersmodel

import (
	"fmt"
	"wms/pkg/utils"
)

type Manufacturer struct {
	ID         		uint64      `gorm:"primaryKey;" json:"id"`
	Name       		string      `gorm:"size:60;not null" json:"name"`
	Country    		string      `gorm:"size:40;not null" json:"country"`
	INN 			string      `gorm:"size:12;not null;uniqueIndex" json:"inn"`
	UrAddress 		string      `gorm:"size:200;not null" json:"urAddress"`
	FactAddress 	string      `gorm:"size:200;not null" json:"factAddress"`
	FIO 			string      `gorm:"size:200;not null" json:"fio"`
	Phone 			string      `gorm:"size:20;not null" json:"phone"`
	Email 			string      `gorm:"size:200;not null" json:"email"`
}

func (Manufacturer) TableName() string {
	return "manufacturers"
}

func (Manufacturer) GetUpdateMap(setRequest any) (map[string]any, error) {
	var err error
	fields :=  map[string]any{}

	fields["name"], err = utils.GetStructFieldByString(setRequest, "Name")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["country"], err = utils.GetStructFieldByString(setRequest, "Country")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["inn"], err = utils.GetStructFieldByString(setRequest, "INN")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["urAddress"], err = utils.GetStructFieldByString(setRequest, "UrAddress")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["factAddress"], err = utils.GetStructFieldByString(setRequest, "FactAddress")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["fio"], err = utils.GetStructFieldByString(setRequest, "FIO")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["phone"], err = utils.GetStructFieldByString(setRequest, "Phone")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["email"], err = utils.GetStructFieldByString(setRequest, "Email")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	return fields, nil
}

func New() *Manufacturer {return &Manufacturer{}}