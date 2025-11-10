package main

import (
	"errors"
)

type Wallet struct {
	balance int
}

var user_balance = map[string]*Wallet{}

func checkUserExist(user string) (bool, *Wallet) {
	data, isUserExist := user_balance[user]
	return isUserExist, data
}

func AddUser(user string) (bool, error) {
	is_user_exist, _ := checkUserExist(user)
	if is_user_exist {
		return false, errors.New("user exist")
	}
	user_balance[user] = &Wallet{balance: 0}

	return true, nil

}

func Deposit(user string, amount int) (bool, error) {
	is_user_exist, data := checkUserExist(user)
	if !is_user_exist {
		return false, errors.New("user does not exisst")
	}

	data.balance += amount
	return true, nil
}

func Withdraw(user string, amount int) (bool, error) {
	is_user_exist, data := checkUserExist(user)
	if is_user_exist && data.balance >= amount {
		data.balance -= amount
		return true, nil
	}

	return false, errors.New("user does not exist")
}

func Balance(user string) int {
	is_user_exist, data := checkUserExist(user)
	if is_user_exist {
		return data.balance
	}
	return 0
}
