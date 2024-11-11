package todos

import "time"

type ToDo struct {
	ID        string
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
