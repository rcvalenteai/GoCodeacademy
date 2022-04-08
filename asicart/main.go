package main

import (
	"fmt"
	"time"
)

func main() {
	printASIC("Dog", dog)
	printASIC("Gopher", gopher)
}

func printASIC(name string, art func()) {
	fmt.Println()
	fmt.Println(name)
	art()
	fmt.Println(time.Now())
}

func dog() {
	fmt.Println("  __      _")
	fmt.Println("o'')}____//")
	fmt.Println(" `_/      )")
	fmt.Println(" (_(_/-(_/ ")
}

func gopher() {
	fmt.Println("    `.-::::::-.`    ")
	fmt.Println(".:-::::::::::::::-:.")
	fmt.Println("`_:::    ::    :::_`")
	fmt.Println(" .:( ^   :: ^   ):. ")
	fmt.Println(" `:::   (..)   :::. ")
	fmt.Println(" `:::::::UU:::::::` ")
	fmt.Println(" .::::::::::::::::. ")
	fmt.Println(" O::::::::::::::::O ")
	fmt.Println(" -::::::::::::::::- ")
	fmt.Println(" `::::::::::::::::` ")
	fmt.Println("  .::::::::::::::.  ")
	fmt.Println("    oO:::::::Oo     ")
}
