package strategy

type Warrior interface {
	SetAttackWay(AttackWay)
	Attack() int
	CriticalAttack() int
}

type warrior struct {
	attackWay AttackWay
}

func NewWarrior(attackway AttackWay) Warrior {
	return &warrior{
		attackWay: attackway,
	}
}

func (w *warrior) SetAttackWay(attackWay AttackWay) {
	w.attackWay = attackWay
}

func (w *warrior) Attack() int {
	return w.attackWay.NormalAttack()
}

func (w *warrior) CriticalAttack() int {
	return w.attackWay.CriticalAttack()
}
