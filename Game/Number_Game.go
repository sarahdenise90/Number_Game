package main

import (
	//"encoding/json"
	"fmt"
	//"io"
	//"net/http"
	//"os"
	"math/rand"
	"time"
)

//Int >= min oder <max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var input int
	fmt.Print("Please enter your number: ")
	fmt.Scan(&input)
	fmt.Println("Your number", input)

	//Returns the now random number
	rand.Seed(time.Now().UnixNano())
	var randomNumber = randomInt(0, 11)
	fmt.Println("The random number is: ", randomNumber)

	//file names lower case

	//Compare the random number with your input
	if randomNumber == input { //equal
		fmt.Println("You got it!")
	} else { //not equal
		fmt.Println("The right number is: ", randomNumber, ", Next try!")
	}
}
