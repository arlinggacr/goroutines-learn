package goroutines_learn

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	// the variable x got used by 1000 goroutines, and each goroutine will increment the value of x by 1 for 100 times, so the final value of x should be 100000,
	//  but because of race condition, the final value of x will be less than 100000
	// technically there is 2 gorountines that is access x simultaneously, and the result will be unpredictable,
	// because the goroutines will read the value of x, increment it, and write it back to x,
	// but if 2 goroutines read the value of x at the same time, they will both increment it and write it back to x, so the final value of x will be less than 100000
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				x = x + 1
			}
		}()

	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", x)

}
