package main

import (
	"fmt"
	"math/rand"
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

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}

	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func getRandomNumber(limit int) int {
	return rand.Intn(limit)
}

func createSpin(reel []string, rows int, cols int) [][]string {
	spin := [][]string{}
	lenth := len(reel)

	for row := 0; row < rows; row++ {
		spin = append(spin, []string{})
		selected := map[int]bool{}

		for col := 0; col < cols; col++ {
			for true {
				randomNumber := getRandomNumber(lenth)

				if !selected[randomNumber] {
					selected[randomNumber] = true
					spin[row] = append(spin[row], reel[randomNumber])
					break
				}
			}
		}
	}

	return spin
}

func printSpin(spin [][]string) {
	print("", "new-line")

	for _, row := range spin {
		for index, symbol := range row {
			print(symbol, "same-line")

			if index != len(row)-1 {
				print(" | ", "same-line")
			}
		}

		print("", "new-line")
	}

	print("", "new-line")
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

	arr := generateSymbolArray(symbols)

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

		spin := createSpin(arr, 3, 3)
		printSpin(spin)

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
