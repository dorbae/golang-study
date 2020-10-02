package main

import (
	"fmt"
	"log"

	"github.com/dorbae/golang-study/banking/accounts"
)

func main() {
	// #2.0 Account + NewAccoutn
	// https://nomadcoders.co/go-for-beginners/lectures/1512
	// account := banking.Account{Owner: "dorbae", Balance: 1000}
	// account := banking.Account{Owner: "dorbae"}
	account := accounts.NewAccount("dorbae")
	fmt.Println(account)

	// #2.1 Methods part One
	// https://nomadcoders.co/go-for-beginners/lectures/1513
	account.Deposit(10)
	fmt.Println(account.Balance())

	// #2.2 Methods part Two
	// https://nomadcoders.co/go-for-beginners/lectures/1514
	// err := account.Withdraw(20)
	err := account.Withdraw(5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())

	// #2.3 Finishig up
	// https://nomadcoders.co/go-for-beginners/lectures/1515
	fmt.Println(account) // Print class
}
