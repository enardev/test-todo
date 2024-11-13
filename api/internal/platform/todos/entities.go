package todos

import "time"

type ToDo struct {
	ID        string    `json:"id" dynamodbav:"id"`
	Title     string    `json:"title" dynamodbav:"title"`
	Completed bool      `json:"completed" dynamodbav:"completed"`
	CreatedAt time.Time `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt time.Time `json:"updated_at" dynamodbav:"updated_at"`
}
