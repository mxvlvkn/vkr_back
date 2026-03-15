package marksmodel

import (
	"fmt"
	"wms/pkg/utils"
)

type Mark struct {
	ID         		uint64      `gorm:"primaryKey;" json:"id"`
	Code       		string      `gorm:"size:200;not null;uniqueIndex" json:"code"`
	NumenclatureID  uint64      `gorm:"not null" json:"numenclatureID"`
}

func (Mark) TableName() string {
	return "marks"
}

func (Mark) GetUpdateMap(setRequest any) (map[string]any, error) {
	var err error
	fields :=  map[string]any{}

	fields["code"], err = utils.GetStructFieldByString(setRequest, "Code")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	fields["numenclatureID"], err = utils.GetStructFieldByString(setRequest, "NumenclatureID")
	if err != nil {
		return nil, fmt.Errorf("GetUpdateMap: %w", err)
	}

	return fields, nil
}

func New() *Mark {return &Mark{}}