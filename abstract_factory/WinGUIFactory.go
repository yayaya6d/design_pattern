package abstract_factory

type WinButton struct {
	text string
}

func (wb *WinButton) Click() {
	println("click win button")
}

func (wb *WinButton) Hover() {
	println("hover on win button")
}

func (wb *WinButton) Text() string {
	return wb.text
}

func (wb *WinButton) Platform() Platform {
	return Windows
}

type WinCheckBox struct {
	toggle bool
}

func (wc *WinCheckBox) SetToggle(b bool) {
	wc.toggle = b
}

func (wc *WinCheckBox) IsToggle() bool {
	return wc.toggle
}

func (wc *WinCheckBox) Platform() Platform {
	return Windows
}

type WinGUIFactory struct{}

func (wf *WinGUIFactory) CreateNewButton(text string) Button {
	return &WinButton{
		text: text,
	}
}

func (wf *WinGUIFactory) CreateNewCheckBox() CheckBox {
	return &WinCheckBox{}
}

func NewWinGUIFactory() GUIFactory {
	return &WinGUIFactory{}
}
