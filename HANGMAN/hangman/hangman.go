package hangman

import ("strings"
		"fmt"
		"time" 
		"math/rand"
		"GoProgram/HANGMAN/history"
		"GoProgram/HANGMAN/player"
		"GoProgram/HANGMAN/config"
		)

var GameHistroy = map[history.Game]string{}
var PlayerGame = map[player.Player]string{}

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
	showToUser := ""
	
	fmt.Println("Guess the Word ", player.Name, " :")
	fmt.Print("Word is ")


	//DONE// TODO: update logic
	/*for i := range word {
		showToUser = strings.Replace(showToUser, string(showToUser[i]), "_", 1)
	}*/

	for i, length := 0, len(word) ; i < length ; i++ {
		showToUser += "_"
	}
	fmt.Println(showToUser)
	

	letter := ""
	var inputLetter []string

	//conf, _ := config.LoadConfiguration("config.json")
	//fmt.Println(conf.Chance)
	//CHANCE, _ := strconv.Atoi(conf.Chance)
	//fmt.Println(CHANCE)
	
	//fetch from configuaration
	var conf, _ = config.LoadConfiguration()
	CHANCE := conf.Chance
	//fmt.Println(CHANCE)
	

	// try to find word
	// TODO: use viper to pick this up from config file
	for chance := CHANCE; chance > 0;  {
		fmt.Println("Enter Letter :")
		fmt.Scan(&letter)
		letter = strings.ToUpper(letter)
		if !checkInputLetter(inputLetter, letter) {
			cnt := 0
			for i := range word {
				if word[i] == letter[0] {
					showToUser = replaceAtIndex(showToUser, letter , i)
					if word == showToUser {
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

		fmt.Println(showToUser)
		
		fmt.Println("Chance left are ", chance)
	}
	return "Dead"
}

func PlayGame(Game int) {
	//DONE// TODO: check default value of seed and why words are not getting randomised
	//Default value of seed is 1
	
	var seedValue int64 = time.Now().UnixNano()
	rand.Seed(seedValue)

	var player player.Player
	wordPool := []string{"GREEN", "BLUE", "RED", "WHITE", "YELLOW", "BLACK"}

	fmt.Println("Enter the User Name : ")
	fmt.Scan(&player.Name)
	PlayerGame[player] = wordPool[rand.Intn(6)]

	Result := hangman(player, PlayerGame[player])

	fmt.Println("You are ", Result , " Word is ", PlayerGame[player])

	//insert history :-----
	var h history.Game

	h.GameNo = Game
	h.PlayerName = player.Name
	h.Word = PlayerGame[player]

	GameHistroy[h] = Result 
}


func CheckHistory() {
	fmt.Println("-------Game History -------")
	for key, value := range GameHistroy {
		fmt.Println(key.GameNo, " ", key.PlayerName, " ", key.Word, " ", value)
	}
	fmt.Println()
}