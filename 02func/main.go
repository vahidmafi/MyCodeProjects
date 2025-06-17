package main

import "fmt"

func main() {
	// fmt.Println("gissele simbana")
	// fmt.Println("9")
	// fmt.Println("I love matcha")
	learn02Variable()
	learn03Inputs()
}

func learn02Variable() {
	var luckyNumber uint8 = 9
	var name string = "gissele simbana"
	var favoriteDrink string = "matcha"

	fmt.Println(name)
	fmt.Println(luckyNumber)
	fmt.Println(favoriteDrink)

	// walrus operator: a new way to declare variables
	// := operator
	myNameIs := "gissele simbana"
	myLuckyNumber := uint8(9)
	myFavoriteDrink := "matcha"

	fmt.Println(myNameIs)
	fmt.Println(myLuckyNumber)
	fmt.Println(myFavoriteDrink)
}

func learn03Inputs() {
	var age int
	fmt.Print("Enter your age: ")
	// ampersand (&) is used to get the memory location address of the variable. So it does not create a new variable, but rather modifies the existing one.
	fmt.Scanln(&age)
	fmt.Println("Your age is:", age)

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Your name is:", name)

	var funFact string
	fmt.Print("Enter a fun fact about yourself: ")
	fmt.Scanln(&funFact)
	fmt.Println("Fun fact about you:", funFact)

	// Scanln reads input from the user until a newline is encountered while Scan reads input until a space is encountered
}
