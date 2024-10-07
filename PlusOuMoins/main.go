package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		lowerBound  int
		higherBound int

		randomNumber int

		try   int
		guess int
	)

	lowerBound = 0
	higherBound = 100

	guess = -1

	try = 1

	randomNumber = getRandomNumber(lowerBound, higherBound)

	// Main loop
	for guess != randomNumber {
		fmt.Printf("#%d> ", try)
		BetterScan(&guess)

		try++

		if guess > randomNumber {
			fmt.Println("Le nombre est plus petit")
		} else {
			fmt.Println("Le nombre est plus grand")
		}
	}

	fmt.Printf("Bravo! Tu as trouvé.e le nombre en %d essai(s)!", try)
}

// Get a random int from lb to hb (inclusive)
func getRandomNumber(lb int, hb int) int {
	return lb + rand.Intn(hb-lb+1)
}

// User input but not fucky
func BetterScan(i any) {
	for ok := true; ok; {
		_, err := fmt.Scan(i)
		if err == nil {
			break
		}
	}
}
