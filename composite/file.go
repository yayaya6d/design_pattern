package composite

type File struct {
	name string
}

func (f *File) Search(name string) bool {
	return f.name == name
}
