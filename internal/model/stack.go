package model

type Stack struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
