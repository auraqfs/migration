package entity

import (
	"github.com/gofrs/uuid"
)

const (
	KelasTableName = "kelas"
)

// KelasModel is a model for entity.Kelas
type Kelas struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Kelas       string    `gorm:"type:varchar(200);not_null" json:"kelas"`
	JumlahSiswa string    `gorm:"type:string(20);null" json:"jumlah_siswa"`
	Auditable
}

func NewKelas(id uuid.UUID, kelas, jumlah_siswa string) *Kelas {
	return &Kelas{
		ID:          id,
		Kelas:       kelas,
		JumlahSiswa: jumlah_siswa,
		Auditable:   NewAuditable(),
	}
}

func (model *Kelas) TableName() string {
	return KelasTableName
}
