package main

import "fmt"

func GetName(balance uint) string {
	name := ""

	print("Welcome to Tim's Casino!", "new-line")
	print("Enter your name: ", "same-line")
	_, err := fmt.Scanf("%s", &name)

	if err != nil {
		return ""
	}

	print(fmt.Sprintf("Hello, %s! ", name), "same-line")
	print(fmt.Sprintf("Your balance is $%d", balance), "new-line")
	return name
}

func GetBet(balance uint) uint {
	var bet uint

	for true {
		print("What's your bet (0 to quit)? $", "same-line")
		fmt.Scan(&bet)

		if bet > balance {
			print(fmt.Sprintf("Insufficient funds. You have only $%d", balance), "new-line")
		} else {
			break
		}
	}

	return bet
}
