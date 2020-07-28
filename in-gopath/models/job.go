package models

// Job is job detail
type Job struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Company string `json:"conpany"`
}
