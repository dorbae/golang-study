package accounts

import "errors"

// Account struct
type Account struct {
	// 대문자 -> Public
	// 소문자 -> Private
	owner   string
	balance int
	// Owner string
	// Balance int

}

// NewAccount creates account(Constructor)
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit method
// 메소드
// func keyword와 func 이름 사이에 receiver를 작성
// receiver는 Strcture의 첫 자 소문자를 써야함
// receiver는 참조 객체
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
// receiver 복사 객체
func (a Account) Balance() int {
	return a.balance
}

// Withdraw money from your account
// Go는 Error를 리턴해줘야함
// Go는 Exception처럼 Event처럼 처리되지 않음
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// return error.Error()
		return errors.New("Can't withdraws your are poor")
	}
	a.balance -= amount
	return nil
}
