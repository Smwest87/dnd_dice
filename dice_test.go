package Dice

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRollDie(t *testing.T) {
	t.Run("Valid Integer -- Success", func(tt *testing.T) {
		randomSeed := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(randomSeed)
		n := randomGenerator.Intn(20)
		result, err := RollDie(n)
		assert.Nil(tt, err)
		assert.Greater(tt, *result, 0)
		assert.Less(tt, *result, 21)

	})

	t.Run("n := 1 ", func(tt *testing.T) {
		n := 1
		result, err := RollDie(n)
		assert.Nil(tt, err)
		assert.Equal(tt, 1, *result)
	})

	t.Run("Failure, n < 1", func(tt *testing.T) {
		n := 0
		expectedErrorMessage := "Die size cannot be 0 or negative"
		result, err := RollDie(n)
		assert.Nil(tt, result)
		assert.Equal(tt, expectedErrorMessage, err.Error())
	})
}

func TestRollStat(t *testing.T) {
	t.Run("Success", func(tt *testing.T) {
		result, err := RollStat()
		assert.Nil(tt, err)
		assert.Greater(tt, *result, 2)
		assert.Less(tt, *result, 19)
	})
}
