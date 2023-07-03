package composite

type Component interface {
	Search(name string) bool
}
