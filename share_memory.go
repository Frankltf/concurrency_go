package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	/*单卡
	 */
	//balance:=80
	//b:=NewBank(NewSimepleAccount(uint8(balance)))
	//fmt.Println("初始化余额", b.Balance())
	//b.Withdraw(uint8(30),"baby")
	//fmt.Println("剩余余额", b.Balance())

	/*
	多卡情况
	 */
	balance:=80
	b:=NewBank(NewConcurrentAccount(uint8(balance)))
	fmt.Println("初始化余额", b.Balance())
	donechan:=make(chan bool)
	go func() {
		b.Withdraw(uint8(30),"daughter")
		donechan<-true
	}()
	go func() {
		b.Withdraw(uint8(10),"son")
		donechan<-true
	}()
	<-donechan
	<-donechan
	fmt.Println("________________")
	fmt.Println("剩余钱",b.Balance())
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
	account.add_some_latency()
	account.balance=balance
}
func (account *SimepleAccount) add_some_latency() {
	<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
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

type ConcurrentAccount struct {
	account *SimepleAccount
	deposit chan uint8
	withdraw chan uint8
	balance chan chan uint8
}

func NewConcurrentAccount(amount uint8)*ConcurrentAccount  {
	acc:=&ConcurrentAccount{
		account:NewSimepleAccount(amount),
		deposit:make(chan uint8),
		withdraw:make(chan uint8),
		balance:make(chan chan uint8),
	}
	acc.listen()
	return acc
}
func (account *ConcurrentAccount)Balance() uint8  {
	ch:=make(chan uint8)
	account.balance<-ch
	return <-ch
}
func (account *ConcurrentAccount)Withdraw(amount uint8)  {
	account.withdraw<- amount
}
func (account *ConcurrentAccount)Deposit(amount uint8)  {
	account.deposit<- amount
}
func (account *ConcurrentAccount)listen()  {
	go func() {
		for {
			select {
			case amt:=<- account.deposit:
				account.account.Deposit(amt)
			case amt:=<-account.withdraw:
				account.account.Withdraw(amt)
			case ch:=<-account.balance:
				ch<-account.account.Balance()
			}
		}
	}()
}
