package model

type Department struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
