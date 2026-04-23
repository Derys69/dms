package dto

type DocumentResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	OwnerID    uint   `json:"owner_id"`
	Department string `json:"department"`
}
