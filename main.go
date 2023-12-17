package main

import (
	"fmt"
	"math/rand"
	"time"
)

func game(you, computer byte) int {
	if you == computer {
		return -1
	}
	if you == 's' && computer == 'p' {
		return 0
	}
	if you == 'p' && computer == 's' {
		return 1
	}
	if you == 's' && computer == 'z' {
		return 1
	}
	if you == 'z' && computer == 's' {
		return 0
	}
	if you == 'p' && computer == 'z' {
		return 0
	}
	if you == 'z' && computer == 'p' {
		return 1
	}
	return -1
}

func main() {
	var n int
	var you, computer byte
	var result int

	rand.Seed(time.Now().UnixNano())

	n = rand.Intn(100)

	if n < 33 {
		computer = 's'
	} else if n > 33 && n < 66 {
		computer = 'p'
	} else {
		computer = 'z'
	}

	fmt.Print("\n\n\n\n\t\t\t\tEnter s for STONE, p for PAPER and z for SCISSOR\n\t\t\t\t\t\t")
	_, err := fmt.Scanf("%c", &you)
	if err != nil {
		return
	}

	result = game(you, computer)

	if result == -1 {
		fmt.Println("\n\n\t\t\t\tGame Draw!")
	} else if result == 1 {
		fmt.Println("\n\n\t\t\t\tWow! You have won the game!")
	} else {
		fmt.Println("\n\n\t\t\t\tOh! You have lost the game!")
	}

	fmt.Printf("\t\t\t\tYou choose: %c and Computer choose: %c \n", you, computer)
}
