package domain

type Experience struct {
	ID      int      `json:"id"`
	Company string   `json:"company"`
	Role    string   `json:"role"`
	Period  string   `json:"period"`
	Duties  []string `json:"duties"`
}
