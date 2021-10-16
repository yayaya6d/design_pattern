package factory_method

type Weapon interface {
	GetName() string
	Attack() int
	AttackDistance() int
}
