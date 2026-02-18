package unitsmodel

type Unit struct {
	ID         		uint64    `gorm:"primaryKey;" json:"id"`
	Name      		string    `gorm:"uniqueIndex;size:40;not null" json:"name"`
	Code   			uint   	  `gorm:"not null" json:"code"`
	Sign       		string    `gorm:"size:10;not null" json:"sign"`
}

func (Unit) TableName() string {
	return "units"
}