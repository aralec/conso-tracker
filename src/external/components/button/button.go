package components

type Button struct {
	Label   string
	Class   string
	OnClick *OnClick
}

type OnClick struct {
	HxTarget  string
	HxTrigger string
	HxSwap    string
	HxUrl     string
	Method    string
}
