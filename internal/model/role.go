package model

type Role struct {
	ID          int          `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
	Users       []User       `json:"users"`
}
