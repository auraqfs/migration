package entity

import (
	"github.com/gofrs/uuid"
)

const (
	GuruTableName = "guru"
)

// GuruModel is a model for entity.Guru
type Guru struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	NamaGuru      string    `gorm:"type:varchar(200);not_null" json:"nama_guru"`
	MataPelajaran string    `gorm:"type:varchar(200);null" json:"mata_pelajaran"`
	Username      string    `gorm:"type:text;null" json:"username"`
	Password      string    `gorm:"type:text;null" json:"password"`
	Auditable
}

func NewGuru(id uuid.UUID, nama_guru, mata_pelajaran, username, password string) *Guru {
	return &Guru{
		ID:            id,
		NamaGuru:      nama_guru,
		MataPelajaran: mata_pelajaran,
		Username:      username,
		Password:      password,
		Auditable:     NewAuditable(),
	}
}

func (model *Guru) TableName() string {
	return GuruTableName
}
