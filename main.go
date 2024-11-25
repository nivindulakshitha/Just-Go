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

func getName() string {
	name := ""

	print("Welcome to Tim's Casino!", "new-line")
	print("Enter your name: ", "same-line")
	_, err := fmt.Scanf("%s", &name)

	if err != nil {
		return ""
	}

	print(fmt.Sprintf("Hello, %s!", name), "new-line")
	return name
}

func main() {
	getName()
}