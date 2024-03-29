package models

type UserRole struct {
	UserId uint64 `gorm:"primary_key;auto_increment:false"`
	RoleId uint64 `gorm:"primary_key;auto_increment:false"`
}

func (UserRole) TableName() string {
	return "users_roles"
}
