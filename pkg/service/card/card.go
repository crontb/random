package card

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateNumbers Returns slice of randomly generated integers from 0 to maxNumber
func GenerateNumbers(totalNumbers int, maxNumber int) ([]int, error) {
	if totalNumbers <= 0 {
		return nil, fmt.Errorf("Wrong totalNumbers: expected more than 0, got %d", totalNumbers)
	}

	result := make([]int, totalNumbers)
	ch := make(chan int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < totalNumbers; i++ {
		go func(ch chan<- int) {
			// do hard random calculation work :-)
			ch <- rand.Intn(maxNumber + 1)
		}(ch)
	}

	for i := 0; i < totalNumbers; i++ {
		result[i] = <-ch
	}
	close(ch)

	return result, nil
}
