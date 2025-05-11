package models

type EmailRequest struct {
	To      []string `json:"to" binding:"required"`
	Cc      []string `json:"cc"` // optional
	Subject string   `json:"subject" binding:"required"`
	Body    string   `json:"body" binding:"required"`
}
