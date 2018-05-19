package main

import "fmt"

func main() {
	balance:=80
	b:=NewBank(NewSimepleAccount(uint8(balance)))
	fmt.Println("初始化余额", b.Balance())
	b.Withdraw(uint8(30),"baby")

	fmt.Println("剩余余额", b.Balance())
}

type Account interface {
	Deposit(uint82 uint8)  //存钱
	Withdraw(uint82 uint8) //取钱
	Balance()uint8  //查看钱
}
type Bank struct {
	account Account
}

func NewBank(account Account)*Bank  {
	return &Bank{account:account}
}
func (bank *Bank)Deposit(amount uint8,name string)  {
	fmt.Println("[+]",amount,name)
	bank.account.Deposit(amount)
}
func (bank *Bank)Withdraw(amount uint8,name string)  {
	fmt.Println("[-]",amount,name)
	bank.account.Withdraw(amount)
}
func (bank *Bank)Balance()uint8  {
	return bank.account.Balance()
}
type SimepleAccount struct {
	balance uint8
}

func NewSimepleAccount(balance uint8)*SimepleAccount  {
	return &SimepleAccount{balance:balance}
}
func (account *SimepleAccount)setBalance(balance uint8)  {
	account.balance=balance
}
func (account *SimepleAccount)Deposit(amount uint8)  {
	account.setBalance(account.balance+amount)
}
func (account *SimepleAccount)Withdraw(amount uint8)  {
	account.setBalance(account.balance-amount)
}
func (account *SimepleAccount)Balance() uint8 {
	return account.balance
}


