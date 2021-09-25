package observer

type Item struct {
	observerList []Observer
	name         string
	avaliable    bool
}

func NewItem(name string, avalible bool) *Item {
	return &Item{
		name:      name,
		avaliable: avalible,
	}
}

func (i *Item) UpdateAvailability(avaliable bool) {
	i.avaliable = avaliable
	i.NotifyAllObservers()
}

func (i *Item) Register(observer Observer) {
	i.observerList = append(i.observerList, observer)
}

func (i *Item) DeRegister(targetObserver Observer) {
	for index, observer := range i.observerList {
		if observer.GetID() == targetObserver.GetID() {
			i.observerList[index] = i.observerList[len(i.observerList)-1]
			i.observerList = i.observerList[:len(i.observerList)-1]
			break
		}
	}
}

func (i *Item) NotifyAllObservers() {
	updateInfo := UpdateInfo{
		i.name,
		i.avaliable,
	}

	for _, observer := range i.observerList {
		observer.Update(updateInfo)
	}
}
