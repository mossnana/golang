package helloworldmodels

// Job is job
type Job struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Company string `json:"conpany"`
}
