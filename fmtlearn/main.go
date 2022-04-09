package main

import "fmt"

func main() {
	printSpacing()
	printFormatting()
	multiTypePrintFormatting()
	stringFormatting()
	sprintfFormatting()
	scanning()
}

func printSpacing() {
	fmt.Println("Let's first see how", "the Println() method works.")
	fmt.Println("Notice that each statement adds a newline for us.")
	fmt.Println("There's also a default space", "between the string arguments.")

	// Add your code below:
	fmt.Print("Print", "is", "different")
	fmt.Print("See?")
}

func printFormatting() {
	animal1 := "cat"
	animal2 := "dog"

	// Add your code below:
	fmt.Printf("Are you a %v or a %v person?", animal1, animal2)
}

func multiTypePrintFormatting() {
	floatExample := 1.75
	// Edit the following Printf for the FIRST step
	fmt.Printf("Working with a %T", floatExample)

	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	// Edit the following Printf for the SECOND step
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp)

	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	// Edit the following Printf for the THIRD step
	fmt.Printf("Each share of Gopher feed is $%.2f", stockPrice)
}

func stringFormatting() {
	step1 := "Breathe in..."
	step2 := "Breathe out..."

	// Add your code below:
	var meditation string = fmt.Sprintln(step1, step2)
	fmt.Print(meditation)
}

func sprintfFormatting() {
	template := "I wish I had a %v."
	pet := "puppy"

	var wish string

	// Add your code below:
	wish = fmt.Sprintf(template, pet)
	fmt.Println(wish)
}

func scanning() {
	fmt.Println("What would you like for lunch?")

	// Add your code below:
	var food string
	scan, err := fmt.Scan(&food)
	if err != nil {
		fmt.Println(scan)
		return
	}

	fmt.Printf("Sure, we can have %v for lunch.", food)
}
