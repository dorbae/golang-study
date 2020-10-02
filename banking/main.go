package main

import (
	"fmt"
	"log"

	"github.com/dorbae/banking/accounts"
)

func main() {
	// account := banking.Account{Owner: "dorbae", Balance: 1000}
	// account := banking.Account{Owner: "dorbae"}
	account := accounts.NewAccount("dorbae")

	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())

}
