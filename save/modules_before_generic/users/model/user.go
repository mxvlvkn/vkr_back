package usersmodel

type User struct {
	ID         		uint64    `gorm:"primaryKey;" json:"id"`
	Login      		string    `gorm:"uniqueIndex;size:40;not null" json:"login"`
	PasswordHash   	string    `gorm:"column:password_hash;size:40;not null" json:"-"`
	Name       		string    `gorm:"size:40;not null" json:"name"`
	Surname    		string    `gorm:"size:40;not null" json:"surname"`
	Patronymic 		string    `gorm:"size:40;not null" json:"patronymic"`
	RoleID       		uint64    `gorm:"not null" json:"role_id"`
}

func (User) TableName() string {
	return "users"
}