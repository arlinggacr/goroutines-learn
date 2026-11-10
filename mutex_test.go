package goroutines_learn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	// so the mutex need to be lock, before the variable got accessed.
	// and we need to unlock the mutex after the variable got accessed, so that other goroutines can access the variable.
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				// the flow is when got the lock
				// the 999 other goroutines will wait until the lock is released,
				//  and then they will get the lock and increment the value of x by 1,
				// so the final value of x will be 100000, which is the expected value
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", x)
}
