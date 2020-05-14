package Dice

import (
	"math/rand"
	"sort"
	"time"
)

func RollDie(n int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := n
	result := rand.Intn(max-1) + min
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

/*func main() {
	dieCount, countErr := strconv.Atoi(os.Args[1])
	if countErr != nil {
		fmt.Println("Please provide two command line arguments, the first being the number of dice to being rolled. The second being the size of the die")
		os.Exit(1)
	}
	dieSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Please supply a second command line argument in the form of an int")
		os.Exit(1)
	}

	rollCount := 0
	rollTotal := 0

	for rollCount < dieCount {
		rollTotal += rollDie(dieSize)
		rollCount++
	}

	fmt.Printf("Total Roll = %v\n", rollTotal)
	fmt.Println("-------------------------")
	fmt.Printf("Strength = %v\n", RollStat())
	fmt.Printf("Dexterity = %v\n", RollStat())
	fmt.Printf("Constitution = %v\n", RollStat())
	fmt.Printf("Wisdom = %v\n", RollStat())
	fmt.Printf("Intelligence = %v\n", RollStat())
	fmt.Printf("Charisma = %v\n", RollStat())
	fmt.Println("-------------------------")

}
*/
