// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine

	balance = 3
	balances <- balance
	for {
		fmt.Println("balance",balance,"sfsd")
		select {
		case amount := <-deposits:
			//println("sdfsfdlksdfkj")
			fmt.Println("============================sdfsfdlksdfkj",)
			balance += amount
			//balances <- balance
		case balances <- balance:
			fmt.Println("bala",balance,time.Now())
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
