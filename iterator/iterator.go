package iterator

type Iterator interface {
	Next() (int, error)
	HasNext() bool
}
