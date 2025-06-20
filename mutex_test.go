package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (a *BankAccount) AddBalance(amount int) {
	a.RWMutex.Lock()
	a.Balance += amount
	a.RWMutex.Unlock()
}

func (a *BankAccount) GetBalance() int {
	a.RWMutex.RLock()
	balance := a.Balance
	a.RWMutex.RUnlock()
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
	fmt.Println("Total Balance", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) UnLock() {
	u.Mutex.Unlock()
}

func (u *UserBalance) Change(amount int) {
	u.Balance += amount
}

func Transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.UnLock()
	user2.UnLock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Dimas",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "Risma",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)

	fmt.Println("User1", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User1", user2.Name, ", Balance ", user2.Balance)
}
