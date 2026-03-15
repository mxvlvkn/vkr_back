package barcodesviewmodel

import (
)

type BarcodeView struct {
	ID         		uint64    `gorm:"primaryKey;" json:"id"`
	Code       		string    `gorm:"size:100;not null" json:"code"`
	Numenclature    string    `gorm:"not null" json:"numenclature"`
	NumenclatureID  uint64    `gorm:"not null" json:"numenclature_id"`
}

func (BarcodeView) TableName() string {
	return "barcodes_view"
}

func (BarcodeView) GetUpdateMap(setRequest any) (map[string]any, error) {
	return map[string]any{}, nil
}

func New() *BarcodeView {return &BarcodeView{}}