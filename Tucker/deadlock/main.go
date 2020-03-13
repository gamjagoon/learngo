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


func (a *account) widthdawo(val int) {
	a.balance -= val
}

func (a *account) deposit(val int) {
	a.balance += val
}

var accounts []*account

func trasfer(sender, receiver int, money int) {
	//! Dead lock example
	// accounts[sender].mutex.Lock()
	// accounts[receiver].mutex.Lock()
	
	accounts[sender].widthdawo(money)
	accounts[receiver].deposit(money)
	
	// accounts[receiver].mutex.Unlock()
	// accounts[sender].mutex.Unlock()

	fmt.Println("transfer",sender,receiver,money)
}

func gettotalbalance() int {
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].balance
	}
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

	go func(){
		for {
			trasfer(0,1,100)
		}
	}()

	go func(){
		for {
			trasfer(1,0,100)
		}
	}()

	for i := 0; i < 1; i++ {
		go gotransfer()
	}

	for {
		time.Sleep(100 * time.Microsecond)
	}
}
