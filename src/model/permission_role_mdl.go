package model

type PermissionRole struct {
	PermissionGUID string     `json:"permission_guid"`
	RoleGUID       string     `json:"role_guid"`
	Permission     Permission `json:"permission" gorm:"foreignKey:PermissionGUID"`
	Role           Role       `json:"role" gorm:"foreignKey:RoleGUID"`
}

func (PermissionRole) TableName() string {
	return "permission_role"
}
