package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("gissele simbana")
	// fmt.Println("9")
	// fmt.Println("I love matcha")

	// learn02Variable()
	// learn03Inputs()
	// learn04Pointers()
	// learn05Bufio()
	learn06Printf()
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
	// bufio package can be used for more advanced input handling issues, such as reading entire lines or handling spaces in input.
	// For example, to read a full line of input including spaces, you can use bufio.NewReader(os.Stdin).ReadString('\n')
	// This will read input until a newline character is encountered, allowing you to capture multi-word
	// howerver bufio needs more memory and is not as efficient for simple inputs.
	// For simple inputs, fmt.Scanln is sufficient and more efficient.
}

func learn04Pointers() {
	var age int = 30
	// * is used to declare a pointer variable, which holds the memory address of another variable.
	// A pointer is a variable that stores the memory address of another variable.
	var agePointer *int
	fmt.Println("agePointer initial value:", agePointer) // This will print <nil> because agePointer is not initialized yet
	agePointer = &age                                    // & is used to get the memory address of the variable

	fmt.Println("Age:", age)
	fmt.Println("Age pointer:", agePointer)
	fmt.Println("Value at age pointer:", *agePointer) // * is used to dereference the pointer to get the value at that memory address

	*agePointer = *agePointer + 31 // changing the value at the memory address
	fmt.Println("Updated Age:", age)

	// why use pointers?
	// 1. To modify the value of a variable in a function without returning it.
	// 2. To save memory by not copying large data structures.
	// 3. To create linked data structures like linked lists, trees, etc.
	// 4. To share data between goroutines in concurrent programming.
	// 5. To avoid unnecessary copying of data when passing large structs or arrays to functions.
	// 6. To implement data structures like maps, slices, and channels that require reference types.
	// 7. To create dynamic memory allocation using the `new` keyword.
	// 8. To implement polymorphism in Go by using interfaces and pointers to structs.
	// 9. To create more efficient algorithms by avoiding unnecessary copies of data.
	// 10. To implement low-level programming techniques like memory management and pointer arithmetic.

}

func learn05Bufio() {
	// bufio package can be used for more advanced input handling issues, such as reading entire lines or handling spaces in input.
	// For example, to read a full line of input including spaces, you can use bufio.NewReader(os.Stdin).ReadString('\n')
	// This will read input until a newline character is encountered, allowing you to capture multi-word
	// however bufio needs more memory and is not as efficient for simple inputs.
	// For simple inputs, fmt.Scanln is sufficient and more efficient.

	// variables can be functions
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name: ")
	// _ is used to ignore the error returned by ReadString. What kind of error can be returned? EOF error if the input is empty or if there is an error reading from the input stream.
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is:", name)

	fmt.Print("Enter your favorite number: ")
	favoriteNumber, _ := reader.ReadString('\n') // ReadString reads until a newline character is encountered
	fmt.Println("Your favorite number is:", favoriteNumber)
	// scanln is more efficient for simple inputs, but bufio is more flexible for complex inputs. thus for numbers, we can use fmt.Scanln
	fmt.Print("Enter your fun fact: ")
	funFact, _ := reader.ReadString('\n')
	fmt.Println("Your fun fact is:", funFact)

	bookReader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your favorite author? ")
	author, _ := bookReader.ReadString('\n')
	fmt.Println("Your favorite author is:", author)

	fmt.Println("What is your book title from", author, "?")
	bookTitle, _ := bookReader.ReadString('\n')
	fmt.Println("Your book title is:", bookTitle)

	fmt.Printf("What is the year of publication for %s?", bookTitle)
	year, _ := bookReader.ReadString('\n')
	fmt.Println("The year of publication is:", year)
}

func learn06Printf() {
	// fmt.Printf is used to format the output
	// %s is used to format a string
	// %d is used to format an integer
	// %f is used to format a float
	// %t is used to format a boolean
	// %v is used to format any value
	// %T is used to print the type of a value

	name := "gissele simbana"
	luckyNumber := 9
	favoriteDrink := "matcha"

	fmt.Printf("My name is %s, my lucky number is %d, and my favorite drink is %s. The data types are %T, %T, and %T.\n", name, luckyNumber, favoriteDrink, name, luckyNumber, favoriteDrink)
}
