package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Description string
}

type Todos []Todo
