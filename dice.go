package Dice

import (
	"errors"
	"math/rand"
	"sort"
	"time"
)

//RollDie -- need to find a way to randomize independent of time
func RollDie(n int) (*int, error) {
	if n < 1 {
		err := errors.New("Die size cannot be 0 or negative")
		return nil, err
	}
	time.Sleep(100 * time.Millisecond)
	randomSeed := rand.NewSource(time.Now().UnixNano())
	randomNumberGenerator := rand.New(randomSeed)
	max := n + 1
	if n == 1 {
		return &n, nil
	}
	result := randomNumberGenerator.Intn(max)
	if result == 0 {
		result = 1
		return &result, nil
	}
	return &result, nil
}

//RollStat -- rolls the stat for players. Roll 4d6 and drop the lowest value
func RollStat() (*int, error) {
	var list []int
	total := 0

	for i := 0; i < 4; i++ {
		roll, err := RollDie(6)
		if err != nil {
			return nil, err
		}
		list = append(list, *roll)
	}

	sort.Ints(list)
	total = list[1] + list[2] + list[3]

	return &total, nil
}
