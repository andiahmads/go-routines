package deadlock_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

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
	//user 1 transfer ammount
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	//user 2 menerima ammount
	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	time.Sleep(1 * time.Second)

	user2.Change(amount)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "andi ",
		Balance: 1000,
	}
	user2 := UserBalance{
		Name:    "desy",
		Balance: 2000,
	}

	//user 1 transfer data ke user 2
	go Transfer(&user1, &user2, 1000)

	//user 2 transfer data ke user 1
	go Transfer(&user2, &user1, 1000)
	time.Sleep(3 * time.Second)
	//case dead lock ketika user 1 transfer ke user 2, dan ketika waktu bersamaan user2 transfer ke user 1

	fmt.Println("User", user1.Name, "Balance", user1.Balance)
	fmt.Println("User", user2.Name, "Balance", user2.Balance)

}
