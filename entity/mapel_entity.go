package entity

import (
	"github.com/gofrs/uuid"
)

const (
	MapelTableName = "mapel"
)

// MapelModel is a model for entity.Mapel
type Mapel struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Nama string    `gorm:"type:varchar(200);not_null" json:"nama"`
	Auditable
}

func NewMapel(id uuid.UUID, nama string) *Mapel {
	return &Mapel{
		ID:        id,
		Nama:      nama,
		Auditable: NewAuditable(),
	}
}

func (model *Mapel) TableName() string {
	return MapelTableName
}
