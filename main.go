//package main
package main

import (
	//import for DiceRoller functionality
	"Dice/DiceRoller"
)

//main function - calls RollDices function by providing diceroller input string
func main() {
	inputString := "2d6+10+2d10"
	DiceRoller.RollDices(inputString)
}