package taskService

import (
	"gorm.io/gorm"
	"time"
)

var task string

type Task struct {
	gorm.Model
	Task   string `json:"task" gorm:"not null"`
	IsDone bool   `json:"is_done" gorm:"default:false"`
}
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
