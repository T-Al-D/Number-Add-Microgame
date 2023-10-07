package main

import (
	"fmt"
	"math/rand"
	"time"

	"bufio"
	"os"
	"strings"
	"strconv"
)

var randomNumber1 int
var randomNumber2 int 

func readInput() string {
	// capture the std.in input (keys), so user can interact
	reader := bufio.NewReader(os.Stdin)
	// once enter is pressed, string is being saved in "name"
	// _ = no need for the error
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}

func gameLogic(input string, result int) {
	// try to convert string to int
	num, err := strconv.Atoi(input)

	if err == nil {
		if num == result {
			fmt.Println("RHIGHT!")
		}	else{
			fmt.Println("WRONG!")
		}
	}else{
		fmt.Println("Error:", err)
	}
}

func main() {

	fmt.Println("Welcome to the Number-ADD game!")
	fmt.Println("Try to add both numbers together and type the correct answer!")
	fmt.Println("Type 'q' to exit the program.")

	// endless for-loop
	for{
		// randomsource has a "seed" to make sure,
		// that the random numbers are more random	
		random := rand.New(rand.NewSource(time.Now().UnixNano()))

		// use the randomsource to generate random numbers until .Intn(number) 
		randomNumber1 = random.Intn(100)
		randomNumber2 = random.Intn(100)
		fmt.Println("NumberOne: ", randomNumber1)
		fmt.Println("NumberTwo: ", randomNumber2)
		var result int = randomNumber1 + randomNumber2

		// take userInput
		var userInput string = readInput()
		
		// if the userInput is "q", loop is breaking
		if userInput == "q"{
			break
		}

		// do something with the input
		gameLogic(userInput, result)

		fmt.Println("Correct answer is: ", randomNumber1 + randomNumber2)
		fmt.Println(" ")
	}
}
