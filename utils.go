package main

import "fmt"

func GetName(balance uint) string {
	name := ""

	Print("Welcome to Tim's Casino!", "new-line")
	Print("Enter your name: ", "same-line")
	_, err := fmt.Scanf("%s", &name)

	if err != nil {
		return ""
	}

	Print(fmt.Sprintf("Hello, %s! ", name), "same-line")
	Print(fmt.Sprintf("Your balance is $%d", balance), "new-line")
	return name
}

func GetBet(balance uint) uint {
	var bet uint

	for true {
		Print("What's your bet (0 to quit)? $", "same-line")
		fmt.Scan(&bet)

		if bet > balance {
			Print(fmt.Sprintf("Insufficient funds. You have only $%d", balance), "new-line")
		} else {
			break
		}
	}

	return bet
}
