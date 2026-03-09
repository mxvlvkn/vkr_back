package numenclaturesviewmodel

import (
)

type NumenclatureView struct {
	ID         			uint64      `gorm:"column:id;primaryKey;" json:"id"`
	Name       			string      `gorm:"column:name;size:400;not null" json:"name"`
	Article 			string      `gorm:"column:article;size:200;not null;uniqueIndex" json:"article"`
	Unit 				string      `gorm:"column:unit;size:40;not null" json:"unit"`
	Manufacturer 		string      `gorm:"column:manufacturer;size:60;not null" json:"manufacturer"`
	UseSerial    		string      `gorm:"column:use_serial;size:5;not null" json:"useSerial"`
	UseMarks 			string      `gorm:"column:use_marks;size:5;not null" json:"useMarks"`
}

func (NumenclatureView) TableName() string {
	return "numenclatures_view"
}

func (NumenclatureView) GetUpdateMap(setRequest any) (map[string]any, error) {
	return map[string]any{}, nil
}

func New() *NumenclatureView {return &NumenclatureView{}}