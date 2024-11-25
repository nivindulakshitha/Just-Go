package main

import (
	"math/rand"
)

func GenerateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}

	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}

	return symbolArr
}

func GetRandomNumber(limit int) int {
	return rand.Intn(limit)
}

func CreateSpin(reel []string, rows int, cols int) [][]string {
	spin := [][]string{}
	lenth := len(reel)

	for row := 0; row < rows; row++ {
		spin = append(spin, []string{})
		selected := map[int]bool{}

		for col := 0; col < cols; col++ {
			for true {
				randomNumber := GetRandomNumber(lenth)

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

func PrintSpin(spin [][]string) {
	Print("", "new-line")

	for _, row := range spin {
		for index, symbol := range row {
			Print(symbol, "same-line")

			if index != len(row)-1 {
				Print(" | ", "same-line")
			}
		}

		Print("", "new-line")
	}

	Print("", "new-line")
}