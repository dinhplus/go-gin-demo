package model

type Permission struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Resource string `json:"resource"`
	Action   string `json:"action"` // view, create, update, delete
}
