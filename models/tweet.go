package models

type Tweet struct {
	Content   string   `json:"content"`
	CreatedAt string   `json:"created_at"`
	CreatedBy uint64   `json:"created_by"`
	Tags      []string `json:"tags"`
}
