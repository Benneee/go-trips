package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Greet and present initial instructions
	fmt.Println("Welcome to the area calculator app")

	var shapeChoice = collectShapeChoice()

	// handle calculations
	switch shapeChoice {
		case "a":
			calculateSquareArea()
		case "b":
			calculateRectangleArea()
		case "c":
			calculateTriangleArea()
		default:
			fmt.Println("You have selected an invalid option")
	}

	var restart string;
	fmt.Printf("Do you wish to compute another shape's area? Enter yes (y) or no (n): ")
	fmt.Scan(&restart)

	if restart == "yes" || restart == "y" {
		main()
	} else {
		fmt.Println("Thank you for using area calculator!")
	}
}

func collectShapeChoice() string {
	var shapeChoice string
	fmt.Println("Please select your desired:\n (a) square \n (b) rectangle \n (c) triangle")
	fmt.Print("Type a, b, or c: ")
	fmt.Scan(&shapeChoice)
	shapeChoice = strings.ToLower(shapeChoice)

	var shapeChoiceCollection = make(map[string]string)
	shapeChoiceCollection["a"] = "square"
	shapeChoiceCollection["b"] = "rectangle"
	shapeChoiceCollection["c"] = "triangle" 
	fmt.Printf("You selected %v\n", shapeChoiceCollection[shapeChoice])

	return shapeChoice
}

func calculateSquareArea() {
	var squareLength int
	fmt.Println("Please, enter the length of the square: ")
	fmt.Scan(&squareLength)
	fmt.Printf("The area of the square is %v square metre\n", math.Pow(float64(squareLength), 2))
}

func calculateRectangleArea() {
	var rectangleLength int
	var rectangleBreadth int
	fmt.Println("Please, enter the length of the rectangle")
	fmt.Scan(&rectangleLength)
	fmt.Println("Please, enter the breadth of the rectangle")
	fmt.Scan(&rectangleBreadth)
	var rectangleArea = float64(rectangleLength) * float64(rectangleBreadth)
	fmt.Printf("The area of the rectangle is %v square metre\n", rectangleArea)
}

func calculateTriangleArea() {
	var triangleBase int
	var triangleHeight int
	fmt.Println("Please, enter the base of the triangle")
	fmt.Scan(&triangleBase)
	fmt.Println("Please, enter the height of the triangle")
	fmt.Scan(&triangleHeight)
	var triangleArea = float64(0.5) * float64(triangleBase) * float64(triangleHeight)
	fmt.Printf("The area of the triangle is %v square metre\n", triangleArea)
}

// ToDo: add validation