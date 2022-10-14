package main
import (
	"fmt"
	"sync"
	"time"
)

type Account struct{
	balance int
	lock sync.Mutex
}

func (a *Account) GetBallance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		fmt.Println("umm")
	} else {
		fmt.Println("withdrawn", v)
		a.balance -= v
	}
}

func main() {
	var acc Account
	acc.balance = 100
	fmt.Println(acc.GetBallance())
	for i:=0;i<12;i++ {
		go acc.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}
