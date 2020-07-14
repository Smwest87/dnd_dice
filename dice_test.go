package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRollDie(t *testing.T) {
	roller := NewRoller()
	t.Run("Valid Integer -- Success", func(tt *testing.T) {
		n := roller.RNG.Intn(20)
		result, err := roller.RollDie(n, roller.RNG)
		assert.Nil(tt, err)
		assert.Greater(tt, result, 0)
		assert.Less(tt, result, 21)

	})

	t.Run("n := 1 ", func(tt *testing.T) {
		n := 1
		result, err := roller.RollDie(n, roller.RNG)
		assert.Nil(tt, err)
		assert.Equal(tt, 1, result)
	})

	t.Run("Failure, n < 1", func(tt *testing.T) {
		n := 0
		expectedErrorMessage := "Die size cannot be 0 or negative"
		result, err := roller.RollDie(n, roller.RNG)
		assert.Equal(tt, -1, result)
		assert.Equal(tt, expectedErrorMessage, err.Error())
	})
}

func TestRollStat(t *testing.T) {
	roller := NewRoller()
	t.Run("Success", func(tt *testing.T) {
		result, err := roller.RollStat()
		assert.Nil(tt, err)
		assert.Greater(tt, result, 2)
		assert.Less(tt, result, 19)
	})
}
