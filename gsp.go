package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
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
	 b:=NewBank(NewLockingAccount(uint8(balance)))
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

type LockAccount struct {
	lock sync.Mutex
	account *SimepleAccount
}

func NewLockingAccount(balance uint8)*LockAccount  {
	return &LockAccount{account:NewSimepleAccount(balance)}
}
func (account *LockAccount)Deposit(amount uint8)  {
	account.lock.Lock()
	defer account.lock.Unlock()
	account.account.Deposit(amount)
}
func (account *LockAccount)Withdraw(amount uint8)  {
	account.lock.Lock()
	defer account.lock.Unlock()
	account.account.Withdraw(amount)
}
func (account *LockAccount)Balance()uint8  {
	account.lock.Lock()
	defer account.lock.Unlock()
	return account.account.Balance()
}

