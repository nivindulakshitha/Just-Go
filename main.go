package main

import "fmt"

func print(text string, method string) {
	switch method {
	case "same-line":
		fmt.Print(text)
		break
	default:
		fmt.Println(text)
	}
}

func getName(balance uint) string {
	name := ""

	print("Welcome to Tim's Casino!", "new-line")
	print("Enter your name: ", "same-line")
	_, err := fmt.Scanf("%s", &name)

	if err != nil {
		return ""
	}

	print(fmt.Sprintf("Hello, %s!", name), "new-line")
	print(fmt.Sprintf("Your balance is $%d", balance), "new-line")
	return name
}

func getBet(balance uint) uint {
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

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}

	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 15,
		"D": 20,
	}

	/* multipliers := map[string]uint {
		"A": 20,
		"B": 15,
		"C": 7,
		"D": 4,
	} */

	arr := generateSymbolArray(symbols)

	balance := uint(200)

	getName(balance)
	print("", "new-line")

	for true {
		bet := getBet(balance)

		if bet == 0 {
			break
		}

		balance -= bet
		print(fmt.Sprintf("* You bet $%d, ", bet), "same-line")
		print(fmt.Sprintf("having balance of $%d currently.", balance), "new-line")
		print("", "new-line")
	}

	print(fmt.Sprintf("\nNow you have $%d", balance), "new-line")
	print("Thank you for join, see you next time!", "same-line")

}
