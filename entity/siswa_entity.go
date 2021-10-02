package entity

import (
	"github.com/gofrs/uuid"
)

const (
	SiswaTableName = "siswa"
)

// SiswaModel is a model for entity.Siswa
type Siswa struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	NamaSiswa string    `gorm:"type:varchar(200);not_null" json:"nama_siswa"`
	Kelas     string    `gorm:"type:varchar(200);null" json:"kelas"`
	Username  string    `gorm:"type:text;null" json:"username"`
	Password  string    `gorm:"type:text;null" json:"password"`

	Auditable
}

func NewSiswa(id uuid.UUID, nama_siswa, kelas, username, password string) *Siswa {
	return &Siswa{
		ID:        id,
		NamaSiswa: nama_siswa,
		Kelas:     kelas,
		Username:  username,
		Password:  password,
		Auditable: NewAuditable(),
	}
}

func (model *Siswa) TableName() string {
	return SiswaTableName
}
