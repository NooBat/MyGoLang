// guess challenges players to guess a random number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"github.com/headfirstgo/keyboard"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	fmt.Println("I have chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	fmt.Printf("Make a guess: ")


	var success bool
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Printf("Make a guess: ")
		input, err := keyboard.
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			success = true
			fmt.Println("Goob job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Printf("Sorry. You didn't guess my number. It was: %d.\n", target)
	}
}
