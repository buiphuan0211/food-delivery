package common

import (
	"time"
)

type SQLModel struct {
	ID int `json:"id" gorm:"column:id"`
	//ObjectID  string     `json:"object_id" gorm:"column:object_id"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
}
