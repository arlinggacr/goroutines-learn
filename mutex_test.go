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

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

// struct like this only use by rw mutex if the struct got readed by more than 1 goroutines in concurent way
// and the struct got written by more than 1 goroutines in concurent way, so we need to use rw mutex to protect the struct

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance:", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "eko",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "budi",
		Balance: 1000000,
	}

	fmt.Println("Initial User1 Balance:", user1.Balance)
	fmt.Println("Initial User2 Balance:", user2.Balance)

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)
	time.Sleep(3 * time.Second)

	// the deadlock will happen because the first goroutine will lock user1 and then try to lock user2,
	// but the second goroutine will lock user2 and then try to lock user1, so both goroutines will wait for each other to release the lock,
	// and the program will be stuck in this state forever.

	fmt.Println("User1 Balance:", user1.Balance)
	fmt.Println("User2 Balance:", user2.Balance)
}
