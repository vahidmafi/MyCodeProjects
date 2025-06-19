package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

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

func learn07Loop() {
	fmt.Println("1 to 20")
	for index := 1; index <= 20; index++ {
		// x := math.Pow(float64(index), 2)
		fmt.Print(index, " ")
	}

	fmt.Println("\nOdd numbers between 1 to 20")
	for index := 1; index <= 20; index = index + 2 {
		// x := math.Pow(float64(index), 2)
		fmt.Print(index, " ")
	}

	var alphabets int8
	fmt.Println("\nA to Z")
	for alphabets = 0; alphabets < 26; alphabets++ {
		char := string('A' + alphabets)
		fmt.Print(char, " ")
	}

	// 2nd solution
	// String data type must always be double-quoted. A rune/char datatype it must reside in single quotation. rune/char will be shown as int32 datatype.
	fmt.Println("\nA to Z")
	for abUpp := 'A'; abUpp <= 'Z'; abUpp++ {
		fmt.Printf("%c ", abUpp)
	}
}

func learn07LoopAlphaChar() {
	word := "Hello Coders"
	for i, r := range word {
		fmt.Printf("Character %d: %c = %d\n", i, r, r)
	}
	fmt.Println(word)
}

func greet(name string) *string {
	temp := "Hello, " + name + " :) "
	return &temp
}

func learn07Function() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	messageTest := greet(name)
	fmt.Println(messageTest, *messageTest)
}

// Create our own specific datatype
type Student struct {
	Name   string
	School string
	Year   uint16
}

type Employee struct {
	name   string
	age    uint8
	job    string
	salary uint16
}

func learn09SrtructStudent() {
	student1 := Student{
		Name:   "Vahid Mafi",
		School: "01 Founders",
		Year:   2025,
	}

	fmt.Printf("%%", student1)
}

func learn09StructEmployee() {
	var pers1, pers2, pers3 Employee
	pers1.name = "Yasin"
	pers1.age = 20
	pers1.job = "IT"
	pers1.salary = 2000

	pers2.name = "Yasin"
	pers2.age = 20
	pers2.job = "IT"
	pers2.salary = 2000

	pers3.name = "Yasin"
	pers3.age = 20
	pers3.job = "IT"
	pers3.salary = 2000

	printPerson(pers1)

}

func printPerson(person Employee) {
	fmt.Printf("Name: %s \n", person.name)
	fmt.Printf("Age: %s \n", person.age)
	fmt.Printf("Job: %s \n", person.job)
	fmt.Printf("Salary: %s \n", person.salary)
}

func main() {
	start := time.Now()

	// learn02Variable()
	// learn03Inputs()
	// learn04Pointers()
	// learn05Bufio()
	// learn06Printf()
	// learn07Loop()
	// learn07LoopAlphaChar()
	// learn07Function()
	learn09SrtructStudent()

	elapsed := time.Since(start)
	fmt.Println("")
	fmt.Printf("Execution took %s", elapsed)
}
