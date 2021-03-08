package domain

import "time"

type Todo struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeteletedAt time.Time `json:"deleted_at" gorm:"index"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
}

type Todos []Todo
