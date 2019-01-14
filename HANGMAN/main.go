package main 

import (
	"fmt"
	"GoProgram/Hangman/hangman"
)

func main() {
	Game := 0
	for {
		fmt.Println("------Hangman Game------")
		option := 0
		fmt.Println("1. Play Game \n2. Game History \n3. Quit")
		fmt.Scan(&option)
		switch option {
			case 1 :
				Game++
				hangman.PlayGame(Game)
			case 2 : 
				hangman.CheckHistory()
			default : 
				return 
		}
	}
}