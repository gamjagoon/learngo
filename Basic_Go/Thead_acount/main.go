package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type account struct {
	balance int
	mutex *sync.Mutex
}

var mutex = &sync.Mutex{}

func (a *account) widthdawo(val int) {
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *account) deposit(val int) {
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

var accounts []*account
var globallock *sync.Mutex

func trasfer(sender, receiver int, money int) {
	globallock.Lock()
	accounts[sender].widthdawo(money)
	accounts[receiver].deposit(money)
	globallock.Unlock()
}

func gettotalbalance() int {
	globallock.Lock()
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].balance
	}
	globallock.Unlock()
	return total
}

func printtotalbalance() {
	fmt.Printf("Total : %d\n", gettotalbalance())
}

func randomtrenfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		if accounts[sender].balance > 0 {
			balance = accounts[sender].balance
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if receiver != sender {
			break
		}
	}

	money := rand.Intn(balance)
	trasfer(sender, receiver, money)
}

func gotransfer() {
	for {
		randomtrenfer()
	}
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &account{balance: 1000,mutex: &sync.Mutex{}})
	}
	globallock = &sync.Mutex{}
	printtotalbalance()

	for i := 0; i < 1; i++ {
		go gotransfer()
	}

	for {
		printtotalbalance()
		time.Sleep(100 * time.Microsecond)
	}
}
