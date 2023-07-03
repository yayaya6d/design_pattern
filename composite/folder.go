package composite

type Folder struct {
	name       string
	components []Component
}

func (f *Folder) Search(name string) bool {
	if f.name == name {
		return true
	}

	for _, c := range f.components {
		if c.Search(name) {
			return true
		}
	}

	return false
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}
