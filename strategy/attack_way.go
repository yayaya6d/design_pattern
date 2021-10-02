package strategy

type AttackWay interface {
	NormalAttack() int
	CriticalAttack() int
}
