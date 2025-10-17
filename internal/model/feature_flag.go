package model

type FeatureFlag struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"unique;not null"`
	Description *string `json:"description"`
	IsEnabled   bool    `json:"is_enabled"`
}
