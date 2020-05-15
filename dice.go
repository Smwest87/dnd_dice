package Dice

import (
	"math/rand"
	"sort"
	"time"
)

//RollDie -- need to find a way to randomize independent of time
func RollDie(n int) int {
	time.Sleep(5)
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := n
	result := rand.Intn(max-min) + min
	return result
}

func RollStat() int {
	var list [4]int
	total := 0
	list[0] = RollDie(6)
	list[1] = RollDie(6)
	list[2] = RollDie(6)
	list[3] = RollDie(6)

	sort.Ints(list[:])

	total = list[1] + list[2] + list[3]

	return total
}
