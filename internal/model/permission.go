package model

type Permission struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	ResourceID  int      `json:"resource_id"`
	Resource    Resource `json:"resource" gorm:"foreignKey:ResourceID"`
	Action      string   `json:"action"`                      // view, create, update, delete
	FeatureFlag string   `json:"feature_flag" gorm:"size:64"` // tên hoặc mã feature flag, rỗng nếu không dùng
	Roles       []Role   `json:"roles" gorm:"many2many:role_permissions;"`
}
