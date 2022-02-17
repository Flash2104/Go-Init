package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("overflow")
	var maxVal = math.MaxInt64
	var minVal = math.MinInt64
	var floatVal = math.MaxFloat64
	fmt.Println(minVal == maxVal+1)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))
	fmt.Println("=============")
	fmt.Println("remainder")
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))
	fmt.Println(negResult)
	fmt.Println(absResult)
	fmt.Println("=============")
	fmt.Println("increment decrement")
	value := 10.2
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value -= 2
	fmt.Println(value)
	value--
	fmt.Println(value)
	fmt.Println("=============")
	fmt.Println("Concatenating")
	greeting := "Hello"
	language := "Go"
	combinedString := greeting + ", " + language
	fmt.Println(combinedString)
	fmt.Println("=============")
	fmt.Println("comparing pointers")
	first := 100
	second := &first
	third := &first
	alpha := 100
	beta := &alpha
	fmt.Println(second == third)
	fmt.Println(second == beta)

	fmt.Println("=============")
	fmt.Println("logical operators")
	maxMph := 50
	passengerCapacity := 4
	airbags := true
	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar
	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)

	fmt.Println("=============")
	fmt.Println("Explicit Type Conversions")
	kayak := 275
	soccerBall := 19.49
	total := float64(kayak) + soccerBall
	fmt.Println(total)
	total1 := kayak + int(math.Floor(soccerBall))
	fmt.Println(total1)
	fmt.Println(math.Floor(soccerBall))
	fmt.Println(math.Ceil(soccerBall))
	fmt.Println(int8(total1))

	fmt.Println("=============")
	fmt.Println("Parsing from Strings")
	// 	ParseBool(str) This function parses a string into a bool value. Recognized string values are "true",
	// "false", "TRUE", "FALSE", "True", "False", "T", "F", "0", and "1".
	// ParseFloat(str,
	// size)
	// This function parses a string into a floating-point value with the specified size, as
	// described in the “Parsing Floating-Point Numbers” section.
	// ParseInt(str,
	// base, size)
	// This function parses a string into an int64 with the specified base and size. Acceptable
	// base values are 2 for binary, 8 for octal, 16 for hex, and 10, as described in the “Parsing
	// Integers” section.
	// ParseUint(str,
	// base, size)
	// This function parses a string into an unsigned integer value with the specified base and
	// size.
	// Atoi(str) This function parses a string into a base 10 int and is equivalent to calling
	// ParseInt(str, 10, 0), as described in the “Using the Integer Convenience Function”
	// section.

	val1 := "true"
	val2 := "false"
	val3 := "not true"
	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)
	fmt.Println("Bool 1", bool1, b1err)
	fmt.Println("Bool 2", bool2, b2err)
	fmt.Println("Bool 3", bool3, b3err)
}
