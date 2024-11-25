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

func main() {
	balance := uint(200)

	getName(balance)
	print("", "new-line")

	bet := getBet(balance)
	print(fmt.Sprintf("You bet $%d", bet), "new-line")

}