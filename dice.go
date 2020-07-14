package dice

import (
	"errors"
	"math/rand"
	"sort"
	"time"
)

//Roller -- RNG will be used to select the result of die rolls
type Roller struct {
	RNG *rand.Rand
}

//MockRoller -- TODO -- how to initialize MockDiceRoller in tests?
type MockRoller interface {
	RollDie(int, *rand.Rand)
	RollStat()
}

//NewRoller -- return new roller struct with a new random source
func NewRoller() *Roller {
	randomSeed := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSeed)
	return &Roller{
		RNG: randomGenerator,
	}
}

//RollDie -- roll single die based off of die size and rng seed
func (r *Roller) RollDie(size int, rng *rand.Rand) (int, error) {
	if size < 1 {
		err := errors.New("Die size cannot be 0 or negative")
		return -1, err
	}

	if size == 1 {
		return 1, nil
	}
	result := r.RNG.Intn(size) + 1

	return result, nil
}

//RollStat -- rolls the stat for players. Roll 4d6 and drop the lowest value
func (r *Roller) RollStat() (int, error) {
	var list []int
	total := 0

	for i := 0; i < 4; i++ {
		roll, err := r.RollDie(6, r.RNG)
		if err != nil {
			return -1, err
		}
		list = append(list, roll)
	}

	sort.Ints(list)
	total = list[1] + list[2] + list[3]

	return total, nil
}
