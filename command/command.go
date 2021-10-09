package command

type SingleCommand interface {
	Execute()
}

type NumCommand interface {
	Execute(int)
}
