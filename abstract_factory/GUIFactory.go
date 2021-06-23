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

func NewGUIFactory(p Platform) GUIFactory {
	switch p {
	case Windows:
		return NewWinGUIFactory()
	case Mac:
		return NewMacGUIFactory()
	}

	println("invalid platform.")
	return nil
}
