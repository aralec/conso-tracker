package modal

// Modal representes a html modal
type Modal struct {
	Title  string `json:"title"`
	Active bool   `json:"active"`
}

func NewModal(title string) *Modal {
	return &Modal{Title: title, Active: true}
}
