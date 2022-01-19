// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	"ex01/bank"
)

func TestWithdraw(t *testing.T) {
	bank.Deposit(100)
	ok := bank.Withdraw(50) // 50, true
	if !ok {
		t.Error("got false, want true")
		return
	}
	ok = bank.Withdraw(50) // 0, true
	if !ok {
		t.Error("got false, want true")
		return
	}
	ok = bank.Withdraw(50) // 0, false
	if ok {
		t.Error("got true, want false")
		return
	}
	if bank.Balance() != 0 {
		t.Errorf("got %d, want 0", bank.Balance())
	}
}

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
