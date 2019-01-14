package main 

import (
	"fmt"
	"GoProgram/HANGMAN/hangman"
)

func main() {
	fmt.Println("Starting Application")
	game := 0
	for {
		fmt.Println("------Hangman Game------")
		option := 0
		fmt.Println("1. Play Game \n2. Game History \n3. Quit")
		fmt.Scan(&option)
		switch option {
			case 1 :
				game++
				hangman.PlayGame(game)
			case 2 : 
				hangman.CheckHistory()
			default : 
				return 
		}
	}
}