package main

import (
	"fmt"
)

func print(text string, method string) {
	switch method {
	case "same-line":
		fmt.Print(text)
		break
	default:
		fmt.Println(text)
	}
}

func checkWin(spin [][]string, multipliers map[string]uint) []int {
	lines := []int{}

	for _, row := range spin {
		symbol := row[0]

		if symbol == row[1] && symbol == row[2] {
			lines = append(lines, int(multipliers[symbol]))
		} else {
			lines = append(lines, 0)
		}
	}

	return lines
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 15,
		"D": 20,
	}

	multipliers := map[string]uint{
		"A": 20,
		"B": 15,
		"C": 7,
		"D": 4,
	}

	arr := GenerateSymbolArray(symbols)

	balance := uint(200)

	GetName(balance)
	print("", "new-line")

	for true {
		bet := GetBet(balance)

		if bet == 0 {
			break
		}

		balance -= bet
		print(fmt.Sprintf("* You bet $%d, ", bet), "same-line")
		print(fmt.Sprintf("having balance of $%d currently.", balance), "new-line")

		spin := CreateSpin(arr, 3, 3)
		PrintSpin(spin)

		winningLines := checkWin(spin, multipliers)

		for _, win := range winningLines {
			if win > 0 {
				balance += uint(win * int(bet))
				print(fmt.Sprintf("You won $%d ($%dx%d), balance is $%d", win*int(bet), bet, win, balance), "new-line")
			}
		}

	}

	print(fmt.Sprintf("\nNow you have $%d", balance), "new-line")
	print("Thank you for join, see you next time!", "same-line")

}
