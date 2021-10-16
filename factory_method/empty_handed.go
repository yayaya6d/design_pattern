package factory_method

type EmptyHanded struct {
}

func (e *EmptyHanded) GetName() string {
	return "empty-handed"
}

func (e *EmptyHanded) Attack() int {
	return 1
}

func (e *EmptyHanded) AttackDistance() int {
	return 1
}
