package model

type Resource struct {
	ID          int          `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	Permissions []Permission `json:"permissions" gorm:"many2many:resource_permissions;"`
}
