package strategy

type Sword struct {
	attackValue int
}

func (s *Sword) NormalAttack() int {
	return s.attackValue
}

func (s *Sword) CriticalAttack() int {
	return s.attackValue * 3
}
