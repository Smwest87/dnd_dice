package dice

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRollDie(t *testing.T) {
	roller := NewRoller()
	t.Run("Valid Integer -- Success", func(tt *testing.T) {
		randomSeed := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(randomSeed)
		n := randomGenerator.Intn(20)
		result, err := roller.RollDie(n, randomGenerator)
		assert.Nil(tt, err)
		assert.Greater(tt, result, 0)
		assert.Less(tt, result, 21)

	})

	t.Run("n := 1 ", func(tt *testing.T) {
		n := 1
		randomSeed := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(randomSeed)
		result, err := roller.RollDie(n, randomGenerator)
		assert.Nil(tt, err)
		assert.Equal(tt, 1, result)
	})

	t.Run("Failure, n < 1", func(tt *testing.T) {
		n := 0
		expectedErrorMessage := "Die size cannot be 0 or negative"
		randomSeed := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(randomSeed)
		result, err := roller.RollDie(n, randomGenerator)
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
