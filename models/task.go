package models

import "time"

// Task !
type Task struct {
	TaskID    int       `json:"ID"`
	Title     string    `json:"Title"`
	DateAdded time.Time `json:"Date Added"`
	IsDone    bool      `json:"IsDone"`
}
