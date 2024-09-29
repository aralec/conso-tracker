package modal

// ModalAction represents the props of footer modal actions.
type ModalAction struct {
	Label      string `json:"label"`
	Classes    string `json:"classes"`
	HxTarget   string `json:"hx-target"`
	HxInclude  string `json:"hx-include"`
	HxTrigger  string `json:"hx-trigger"`
	HxSwap     string `json:"hx-swap"`
	HxURL      string `json:"hx-url"`
	HTTPMethod string `json:"http-method"`
}
