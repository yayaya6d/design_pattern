package abstract_factory

type MacButton struct {
	text string
}

func (mb *MacButton) Click() {
	println("click mac button")
}

func (mb *MacButton) Hover() {
	println("hover on mac button")
}

func (mb *MacButton) Text() string {
	return mb.text
}

func (mb *MacButton) Platform() Platform {
	return Mac
}

type MacCheckBox struct {
	toggle bool
}

func (mc *MacCheckBox) SetToggle(b bool) {
	mc.toggle = b
}

func (mc *MacCheckBox) IsToggle() bool {
	return mc.toggle
}

func (wc *MacCheckBox) Platform() Platform {
	return Mac
}

type MacGUIFactory struct{}

func (mf *MacGUIFactory) CreateNewButton(text string) Button {
	return &MacButton{
		text: text,
	}
}

func (mf *MacGUIFactory) CreateNewCheckBox() CheckBox {
	return &MacCheckBox{}
}

func NewMacGUIFactory() GUIFactory {
	return &MacGUIFactory{}
}
