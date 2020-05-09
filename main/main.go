// https://www.hackerrank.com/challenges/truck-tour/problem

package main

import "fmt"

func truckTour(petrolpumps [][]int) int {
	var petrolLeftData = make([]int, len(petrolpumps))
	var positivePetrolLeft = make([]int, 0)

	// list the petrol left for each trip
	for i := 0; i < len(petrolpumps); i++ {
		var left = petrolpumps[i][0] - petrolpumps[i][1]
		petrolLeftData[i] = left

		if left >= 0 {
			// store the station index where the petrol is sufficient to move to the next station
			// trip must be started from these stations
			positivePetrolLeft = append(positivePetrolLeft, i)
		}
	}

	// greedily check all the positive petrol stations
	for i := 0; i < len(positivePetrolLeft); i++ {
		var petrolLeft int

		for j := positivePetrolLeft[i]; positivePetrolLeft[i]-j != 1; j++ {
			petrolLeft += petrolLeftData[j]

			// if the petrol left is negative, straight away break the loop
			// the trip is not possible
			if petrolLeft < 0 {
				break
			}

			if j == len(petrolLeftData)-1 {
				// rotate the index
				j = -1
			}
		}

		// a full cycle has been completed and the petrol left is >= 0 (enough or more than enough)
		// trip is possible from this station
		if petrolLeft >= 0 {
			return positivePetrolLeft[i]
		}
	}

	return -1
}

func main() {
	fmt.Println(truckTour([][]int{{1, 5}, {10, 3}, {3, 4}}))
}
