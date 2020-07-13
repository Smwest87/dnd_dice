package dice

import (
	"errors"
	"math/rand"
	"sort"
	"time"
)

type Roller struct {
}

type mockRoller interface {
	RollDie()
	RollStat()
}

func NewRoller() *Roller {
	roller := Roller{}
	return &roller
}

//RollDie -- need to find a way to randomize independent of time
func (r *Roller) RollDie(size int, rng *rand.Rand) (int, error) {
	if size < 1 {
		err := errors.New("Die size cannot be 0 or negative")
		return -1, err
	}

	if size == 1 {
		return 1, nil
	}
	result := rng.Intn(size) + 1

	return result, nil
}

//RollStat -- rolls the stat for players. Roll 4d6 and drop the lowest value
func (r *Roller) RollStat() (int, error) {
	var list []int
	total := 0

	rngSource := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(rngSource)

	for i := 0; i < 4; i++ {
		roll, err := r.RollDie(6, rng)
		if err != nil {
			return -1, err
		}
		list = append(list, roll)
	}

	sort.Ints(list)
	total = list[1] + list[2] + list[3]

	return total, nil
}
