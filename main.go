package main
import "fmt"

func getName() {
	name := "Nivindu"
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s!", name)
}

func main() {
	getName()
}