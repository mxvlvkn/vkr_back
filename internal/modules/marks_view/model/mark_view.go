package marksviewmodel

import (
)

type MarkView struct {
	ID         		uint64    `gorm:"primaryKey;" json:"id"`
	Code       		string    `gorm:"size:200;not null" json:"code"`
	Numenclature    string    `gorm:"not null" json:"numenclature"`
	NumenclatureID  uint64    `gorm:"not null" json:"numenclature_id"`
}

func (MarkView) TableName() string {
	return "marks_view"
}

func (MarkView) GetUpdateMap(setRequest any) (map[string]any, error) {
	return map[string]any{}, nil
}

func New() *MarkView {return &MarkView{}}