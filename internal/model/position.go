package model

type Position struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
