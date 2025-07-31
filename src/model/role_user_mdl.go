package model

type RoleUser struct {
	RoleGUID string `json:"role_guid"`
	UserGUID string `json:"user_guid"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleGUID"`
	User     User   `json:"user" gorm:"foreignKey:UserGUID"`
}

func (RoleUser) TableName() string {
	return "role_user"
}
