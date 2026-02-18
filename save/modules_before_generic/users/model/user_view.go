package usersmodel

type UserView struct {
	ID         		uint64    `gorm:"primaryKey;" json:"id"`
	Login      		string    `gorm:"uniqueIndex;size:40;not null" json:"login"`
	Name       		string    `gorm:"size:40;not null" json:"name"`
	Surname    		string    `gorm:"size:40;not null" json:"surname"`
	Patronymic 		string    `gorm:"size:40;not null" json:"patronymic"`
	Role       		string    `gorm:"size:40;not null" json:"role"`
}

func (UserView) TableName() string {
	return "users_view"
}