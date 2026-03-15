package numenclaturesmodel

import (
	"fmt"
	"wms/pkg/utils"
)

type Numenclature struct {
	ID         			uint64      `gorm:"primaryKey;" json:"id"`
	Name       			string      `gorm:"size:400;not null" json:"name"`
	Article 			string      `gorm:"size:200;not null;uniqueIndex" json:"article"`
	UseSerial    		bool        `gorm:"not null" json:"useSerial"`
	UseMarks 			bool        `gorm:"not null" json:"useMarks"`
	UnitID 				uint64      `gorm:"not null" json:"unitID"`
	ManufacturerID 		uint64      `gorm:"not null" json:"manufacturerID"`
	ImageURL 			string      `gorm:"not null" json:"ImageURL"`
}


func (Numenclature) TableName() string {
	return "numenclatures"
}

func (Numenclature) GetUpdateMap(setRequest any) (map[string]any, error) {
	var err error
	fields :=  map[string]any{}

	fields["name"], err = utils.GetStructFieldByString(setRequest, "Name")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["useSerial"], err = utils.GetStructFieldByString(setRequest, "UseSerial")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["unitID"], err = utils.GetStructFieldByString(setRequest, "UnitID")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["manufacturerID"], err = utils.GetStructFieldByString(setRequest, "ManufacturerID")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["article"], err = utils.GetStructFieldByString(setRequest, "Article")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["useMarks"], err = utils.GetStructFieldByString(setRequest, "UseMarks")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["imageURL"], err = utils.GetStructFieldByString(setRequest, "ImageURL")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	return fields, nil
}

func New() *Numenclature {return &Numenclature{}}