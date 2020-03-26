package card

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// GenerateNumbers Returns slice of randomly generated integers from 0 to maxNumber
func GenerateNumbers(totalNumbers int, maxNumber int) ([]int, error) {
	if totalNumbers <= 0 {
		return nil, fmt.Errorf("Wrong totalNumbers: expected more than 0, got %d", totalNumbers)
	}

	ch := make(chan int, totalNumbers)
	wg := &sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < totalNumbers; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, ch chan<- int) {
			defer wg.Done()
			// do hard random calculation work :-)
			ch <- rand.Intn(maxNumber + 1)
		}(wg, ch)
	}

	wg.Wait()
	close(ch)

	result := make([]int, 0)
	for number := range ch {
		result = append(result, number)
	}

	return result, nil
}
