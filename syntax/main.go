package main

import "fmt"

func main() {
	literal()
	constants()
	variables()
	setVariables()
	stringBasics()
	diffTypes()
	inferType()
	defaultIntType()
	receipt()
	multiVariable()
}

func literal() {
	fmt.Println(2235 * 1231)
}

func constants() {
	const earthsGravity = 9.80665
	// Create the constant earthsGravity
	// here and set it to 9.80665

	// Here's where we print out the gravity:
	fmt.Println(earthsGravity)
}

func variables() {
	var jellybeanCounter int8
	fmt.Println(jellybeanCounter)
}

func setVariables() {
	var numOfFlavors int
	// Assign a value for numOfFlavors below:
	numOfFlavors = 57

	fmt.Println(numOfFlavors)

	// Declare flavorScale below:
	var flavorScale float32 = 5.8

	fmt.Println(flavorScale)
}

func stringBasics() {
	// Define a string variable
	// called favoriteSnack
	var favoriteSnack string

	// Assign a value to
	// favoriteSnack
	favoriteSnack = "Beef Jerky"

	// Print out the message
	fmt.Println("My favorite snack is " + favoriteSnack)
	// "My favorite snack is "
	// followed by the value in
	// favoriteSnack

}

func diffTypes() {
	// Create three variables
	// emptyInt an int8
	var emptyInt int8

	// emptyFloat a float32
	var emptyFloat float32

	// and emptyString a string
	var emptyString string

	// Finally, print them all out
	fmt.Println(emptyInt)
	fmt.Println(emptyFloat)
	fmt.Println(emptyString)
	fmt.Println(emptyInt, emptyFloat, emptyString)
}

func inferType() {
	// Define daysOnVacation using := below:
	daysOnVacation := 3

	// Define hoursInDay using var and = below:
	var hoursInDay = 24

	fmt.Println("You have spent", daysOnVacation*hoursInDay, "hours on vacation.")
}

func defaultIntType() {
	// Define cupsOfCoffeeConsumed here
	var cupsOfCoffeeConsumed int

	// Give a value to cupsOfCoffeeConsumed
	cupsOfCoffeeConsumed = 3

	// Print out cupsOfCoffeeConsumed
	fmt.Println(cupsOfCoffeeConsumed)
}

func receipt() {
	coolSneakers := 65.99
	niceNecklace := 45.50

	// Define taxCalculation here
	var taxCalculation float64

	// Add coolSneakers to taxCalculation
	taxCalculation += coolSneakers

	// Add niceNecklace to taxCalculation
	taxCalculation += niceNecklace

	// Compute the NYC sales tax
	// 8.875% of the purchase here:

	taxCalculation *= 0.08875

	// Uncomment this line for a receipt!
	fmt.Println("Purchase of", coolSneakers+niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers+niceNecklace+taxCalculation)
}

func multiVariable() {
	// Define magicNum and powerLevel below:
	var magicNum, powerLevel int32
	magicNum = 2048
	powerLevel = 9001

	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)

	// Define amount and unit below:
	amount, unit := 10, "doll hairs"

	fmt.Println(amount, unit, ", that's expensive...")
}
