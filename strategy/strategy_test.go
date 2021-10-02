package strategy

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type strategyTestSuite struct {
	suite.Suite
}

func TestStrategyTestSuite(t *testing.T) {
	suite.Run(t, new(strategyTestSuite))
}

func (s *strategyTestSuite) TestSwordAttack_GivenWarriorSwordAttack_AttackBySword() {
	// arrange
	sword := &Sword{
		attackValue: 4,
	}
	warrior := NewWarrior(sword)

	// act
	normalAttackVal := warrior.Attack()
	criticalAttack := warrior.CriticalAttack()

	// assert
	s.Equal(4, normalAttackVal)
	s.Equal(12, criticalAttack)
}

func (s *strategyTestSuite) TestMagicAttack_GivenWarriorMagicAttack_AttackByMagic() {
	// arrange
	magic := &Magic{
		magicAattackValue: 4,
	}
	warrior := NewWarrior(magic)

	// act
	normalAttackVal := warrior.Attack()
	criticalAttack := warrior.CriticalAttack()

	// assert
	s.Equal(2, normalAttackVal)
	s.Equal(20, criticalAttack)
}

func (s *strategyTestSuite) TestSetAttackWay_GivenWarriorSwordAttackAndSetMagicAttackToIt_AttackByMagic() {
	// arrange
	sword := &Sword{
		attackValue: 4,
	}
	warrior := NewWarrior(sword)

	// act
	magic := &Magic{
		magicAattackValue: 4,
	}
	warrior.SetAttackWay(magic)
	normalAttackVal := warrior.Attack()
	criticalAttack := warrior.CriticalAttack()

	// assert
	s.Equal(2, normalAttackVal)
	s.Equal(20, criticalAttack)
}
