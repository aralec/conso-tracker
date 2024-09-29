package modal

// Modal representes a html modal
type Modal struct {
	Title         string       `json:"title"`
	Active        bool         `json:"active"`
	ConfirmAction *ModalAction `json:"confirm_action"`
	CancelAction  *ModalAction `json:"cancel_action"`
}
