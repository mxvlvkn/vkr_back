package rolesmodel

type Role struct {
	ID         		uint64    `gorm:"primaryKey;" json:"id"`
	Name      		string    `gorm:"uniqueIndex;size:40;not null" json:"name"`
}

func (Role) TableName() string {
	return "roles"
}