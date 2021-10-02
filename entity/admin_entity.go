package entity

import (
	"github.com/gofrs/uuid"
)

const (
	AdminTableName = "admin"
)

// AdminModel is a model for entity.Admin
type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Nama     string    `gorm:"type:varchar(200);not_null" json:"nama"`
	Username string    `gorm:"type:text;null" json:"username"`
	Password string    `gorm:"type:text;null" json:"password"`
	Auditable
}

func NewAdmin(id uuid.UUID, nama, username, password string) *Admin {
	return &Admin{
		ID:        id,
		Nama:      nama,
		Username:  username,
		Password:  password,
		Auditable: NewAuditable(),
	}
}

func (model *Admin) TableName() string {
	return AdminTableName
}
