package manga

type Server struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Url          string `json:"url"`
	Search       bool   `json:"search"`
	StatusActive bool   `json:"status_active"`
}
