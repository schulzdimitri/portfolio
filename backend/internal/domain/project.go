package domain

type Project struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Github      string   `json:"github"`
	Tags        []string `json:"tags"`
}
