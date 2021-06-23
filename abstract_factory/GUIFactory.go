package abstract_factory

type Button interface {
	Click()
	Hover()
	Text() string
	Platform() Platform
}

type CheckBox interface {
	SetToggle(bool)
	IsToggle() bool
	Platform() Platform
}

type GUIFactory interface {
	CreateNewButton(string) Button
	CreateNewCheckBox() CheckBox
}
