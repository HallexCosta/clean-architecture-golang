package entities

import "time"

// User ...
type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreateAt  time.Time
	UpdatedAt time.Time
}
