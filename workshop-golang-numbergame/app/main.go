package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	rand.Seed(time.Now().UnixNano())

	var (
		name string = "tojima"
		answer int = rand.Intn(10) + 1
		count int = 0
		guess int 
	)

	for {
		fmt.Printf("Hello World %v \n", name)
		fmt.Print("Your guess? ")
		fmt.Scanf("%v", &guess)

		count++

		if (answer == guess) {
			fmt.Printf("Bingo! It took %v guesses! \n", count)

			break
		} else if (answer > guess){
			fmt.Println("The answer is higher")
		}else {
			fmt.Println("The answer is lower")
		}

		fmt.Println("-------------------")
	}


	
}