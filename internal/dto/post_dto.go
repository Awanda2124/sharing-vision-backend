package dto

type PostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

type PostResponse struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Content  string `json:"content"`
    Category string `json:"category"`
    Status   string `json:"status"`
}
