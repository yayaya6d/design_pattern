package strategy

type Magic struct {
	magicAattackValue int
}

func (s *Magic) NormalAttack() int {
	return s.magicAattackValue / 2
}

func (s *Magic) CriticalAttack() int {
	return s.magicAattackValue * 5
}
