package factory_method

type Sword struct {
}

func (s *Sword) GetName() string {
	return "sword"
}

func (s *Sword) Attack() int {
	return 5
}

func (s *Sword) AttackDistance() int {
	return 2
}
