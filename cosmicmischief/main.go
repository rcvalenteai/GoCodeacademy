package main

import (
	"fmt"
	"time"
)

const baseBookCost float32 = 8.25
const pageNumberBookCostMultiplier float32 = 0.02
const gradeBookCostMultiplier float32 = 0.50
const ageBookCostMultiplier float32 = -0.05

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade, cost float32

	// book 1
	title = "Mr.GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5
	cost = getBookCost(year, grade, pageNumber)

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher,
		"year:", year, "pages:", pageNumber, "grade:", grade, "cost:", cost)

	// book 2
	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paper Clips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0
	cost = getBookCost(year, grade, pageNumber)

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher,
		"year:", year, "pages:", pageNumber, "grade:", grade, "cost:", cost)
}

func getBookCost(year int, grade float32, pageNumber int) float32 {
	var age int = time.Now().Year() - year
	var cost float32 = baseBookCost +
		(float32(pageNumber) * pageNumberBookCostMultiplier) +
		(grade * gradeBookCostMultiplier) +
		(float32(age) * ageBookCostMultiplier)
	return cost
}
