// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdrawResult struct {
	amount  int
	success chan bool // 取引が成功したか、残高不足で失敗したか
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan *withdrawResult)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	success := make(chan bool)
	withdraws <- &withdrawResult{amount, success}
	return <-success
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraws:
			if balance < withdraw.amount {
				withdraw.success <- false
			} else {
				balance -= withdraw.amount
				withdraw.success <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-