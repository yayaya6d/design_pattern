package observer

type UpdateInfo struct {
	ItemName  string
	Avaliable bool
}

type Observer interface {
	Update(UpdateInfo)
	GetID() string
}
