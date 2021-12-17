package models

import (
	"strings"

	"gorm.io/gorm"
)

// Question model.
type Question struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Question string `json:"question" valid:"required"`
	Answer   string `json:"answer" valid:"required,length(8|50)"`

	CreatedAt int `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt int `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt int `json:"deleted_at,omitempty"`
}

func (Question) TableName() string {
	return "users"
}

func (q *Question) setData() {

	q.Answer = strings.ToLower(q.Answer)
}

func (m *Question) BeforeCreate(tx *gorm.DB) (err error) {
	m.setData()
	if m.DeletedAt == 0 {
		tx.Statement.Omits = append(tx.Statement.Omits, "deleted_at")
	}
	return
}

func (m *Question) BeforeUpdate(tx *gorm.DB) (err error) {
	m.setData()
	return
}
