package main

import "fmt"

func main() {
	// set up necessary variables
	user_one := "akbar-1"
	user_two := "akbar-2"
	unknown_user := "unknown"

	// 1. add new user 1
	AddUser(user_one)

	// 2. deposit balance to user 1
	Deposit(user_one, 10)

	// 3. add new user 2
	AddUser(user_two)

	// 4. withdraw balance user 1
	Withdraw(user_one, 5)

	// 5. deposit balance user 2
	Deposit(user_two, 50)

	// 6. deposit to unknown user
	isDepositSuccess, err := Deposit(unknown_user, 50)
	if !isDepositSuccess {
		fmt.Println("Error: ", err)
	}

	// 7. withdraw to unknown user
	if isSuccess, err := Withdraw(unknown_user, 50); !isSuccess {
		fmt.Println("Error", err)
	}

	// 8. withdraw with amount is greater than balance
	if isSuccess, err := Withdraw(user_one, 500); !isSuccess {
		fmt.Println("Erro: ", err)
	}

	// 9. check balance user 1
	balance_user_one := Balance(user_one)
	fmt.Printf("User balance: %d \n", balance_user_one)

	// 10. check balance user 2
	balance_user_two := Balance(user_two)
	fmt.Printf("User Balance two: %d", balance_user_two)
}
