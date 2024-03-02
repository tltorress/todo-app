package task

import "time"

type Task struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedDate *time.Time `json:"created_date"`
	LastUpdated *time.Time `json:"last_updated"`
	FinishDate  *time.Time `json:"finish_date"`
}
