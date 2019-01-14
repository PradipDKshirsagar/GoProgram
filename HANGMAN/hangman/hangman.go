package hangman

import ("strings"
		"fmt"
		"time" 
		"math/rand"
		"GoProgram/Hangman/history"
		"GoProgram/Hangman/player"
		)

var h = map[history.Game]string{}
var m = map[player.Player]string{}

func checkInputLetter(inputLetter []string, letter string) bool {
	for _, x :=range inputLetter {
		if x == letter {
			return true
		}
	}
	return false 
}

func replaceAtIndex(str string, letter string, index int) string {
    return str[:index] + string(letter) + str[index+1:]
}

func hangman(player player.Player, word string) (string){
	priUser := word 
	
	fmt.Println("Guess the Word ", player.Name, " :")
	fmt.Print("Word is ")

	// TODO: update logic
	for i := range word {
		priUser = strings.Replace(priUser, string(priUser[i]), "_", 1)
	}
	fmt.Println(priUser)
	
	letter := ""
	var inputLetter []string

	//try to find word
	// TODO: use viper to pick this up from config file
	for chance := 6; chance > 0;  {
		fmt.Println("Enter Letter :")
		fmt.Scan(&letter)
		letter = strings.ToUpper(letter)
		//fmt.Println(letter)
		if !checkInputLetter(inputLetter, letter) {
			cnt := 0
			for i := range word {
				if word[i] == letter[0] {
					priUser = replaceAtIndex(priUser, letter , i)
					if word == priUser {
						return "Alive"
					}
				}else {
					cnt++
				}
			}
			if cnt == len(word) {
				chance--
			}
			inputLetter = append(inputLetter, letter)
		}else{
			fmt.Println("Already entered letter  .....  ")	
		}	

		fmt.Println(priUser)
		
		fmt.Println("Chance left are ", chance)
	}
	return "Dead"
}


func PlayGame(Game int) {
	// TODO: check default value of seed and why words are not getting randomised
	var seedValue int64 = time.Now().UnixNano()
	rand.Seed(seedValue)
	var player player.Player
	wordPool := []string{"GREEN", "BLUE", "RED", "WHITE", "YELLOW", "BLACK"}

	fmt.Println("Enter the User Name : ")
	fmt.Scan(&player.Name)
	m[player] = wordPool[rand.Intn(5)]

	Result := hangman(player, m[player])

	fmt.Println("You are ", Result , " Word is ", m[player])

	//insert history :--
	var playerH history.Game

	playerH.GameNo = Game
	playerH.PlayerName = player.Name
	playerH.Word = m[player]

	h[playerH] = Result 
}


func CheckHistory() {
	fmt.Println("--------Game History ------- ")
	for key, value := range h {
		fmt.Println(key.GameNo, " ", key.PlayerName, " ", key.Word, " ", value)
	}
	fmt.Println()
}