package main

import (
	"log"
	"sync"
	"time"
)

type Account struct {
	balance int
	lock    sync.Mutex
}

func main() {
	var acc Account
	acc.balance = 100
	log.Println("Balance :", acc.GetBalance())
	for i := 0; i < 12; i++ {
		go acc.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(value int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if value > a.balance {
		log.Println("Not enough money in account")
	} else {
		log.Printf("%d Withdrawn : Balance :%d\n", value, a.balance)
		a.balance -= value
	}
}

func isEven(num int) bool {
	return num%2 == 0
}

func getEven() {
	num := 4
	var mutex sync.Mutex

	go func() {
		mutex.Lock()
		defer mutex.Unlock()

		numIsEven := isEven(num)
		time.Sleep(5 * time.Millisecond)
		if numIsEven {
			log.Println(num, "is even")
			return
		}
		log.Println(num, "is odd")
	}()

	go func() {
		mutex.Lock()
		num++
		mutex.Unlock()
	}()

	time.Sleep(time.Second)
}
