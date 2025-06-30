package model

type User struct {
	ID           int        `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	DepartmentID int        `json:"department_id"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
	PositionID   int        `json:"position_id"`
	Position     Position   `json:"position" gorm:"foreignKey:PositionID"`
	StackID      int        `json:"stack_id"`
	Stack        Stack      `json:"stack" gorm:"foreignKey:StackID"`
	RoleID       int        `json:"role_id"`
	Role         Role       `json:"role" gorm:"foreignKey:RoleID"`
}
