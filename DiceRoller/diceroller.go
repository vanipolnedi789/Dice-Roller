//package DiceRoller - contains all dice roller related functionalities
package DiceRoller

import (
	//import for logging
	"log"
	"math/rand"
	"time"
	"strings"
	"strconv"
)

//isValidDiceSize - validates dice size (6/10/20)
// Input - dice size 
// Returns :
// 		 1. true - if dice size is 6/10/20
//		 2. false - otherwise
func isValidDiceSize(diceSize int) bool {
    switch diceSize {
    case
        6,
        10,
        20 :
        return true
    }
    return false
}

//duplicateDiceCheck - checks whether any duplicate dice size is provided
// Input - array of dice sizes
// Returns :
// 		1 . duplicate element count if any duplicate found
//      2 . -1 if no duplicate element is found
func duplicateDiceCheck(arr []int) int{
	visited := make(map[int]bool, 0)
	for i:=0; i<len(arr); i++{
	   if visited[arr[i]] == true{
		  return arr[i]
	   } else {
		  visited[arr[i]] = true
	   }
	}
	return -1
}

//sumRollValue - add roll values and writes the sum value into channel
// Input : Consumes rolls arrays (Ex : []string{"3d6","5d10","5"})
// splits roll & dice values  and calls roll function to get the roll values
// adds all roll values and writes the total into ResultChannel
func sumRollValue(rolls []string,ResultChannel chan int) {
	total := 0
	for _, rollValue := range rolls {
		if strings.Contains(rollValue,"d") {
			sub := strings.Split(rollValue,"d")
			rollsCount, _ := strconv.Atoi(sub[0])
			if rollsCount == 0 {
				rollsCount = 1
			}
			diceSize, _ := strconv.Atoi(sub[1])
			total += roll(rollsCount,diceSize)
		} else {
			intVal,_ := strconv.Atoi(rollValue)
			log.Println("adding value ",rollValue," to the roll sum")
			total += intVal
		}
	}
	ResultChannel <- total
}

//splitDices - splits diceRoller using delimiter "+"
func splitDices(diceString string) []string{
	return strings.Split(diceString,"+")
}

//isValidDiceString - verifies diceRoller input string
// This function verifies below cases :
// 		1. Verfies "d" is present in string or not
// 		2. Verifes the dice size (valid input : 6,10,20)
// 		3. Checks any duplicate dice size
// Return Value :
// 		1. returns true if input string is valid 
// 		2. returns false if input string fails any one of above cases 
func isValidDiceString(diceInput string) bool{
	var duplicateDiceArray []int
	if !strings.Contains(diceInput,"d") {
		log.Fatal("dice(d) is mandatory in input string")
		return false
	} else {
		rollsArray := splitDices(diceInput)
		for _,roll := range rollsArray {
			if strings.Contains(roll,"d") {
				rollArray := strings.Split(roll,"d")
				diceSize, _ := strconv.Atoi(rollArray[1])
				if !isValidDiceSize(diceSize) {
					log.Fatal("invalid dice size")
					return false
				}
				duplicateDiceArray = append(duplicateDiceArray,diceSize)
			}
		}
		if duplicateDiceCheck(duplicateDiceArray) > 0 {
			log.Fatal("duplicate dice provided")
			return false
		}
		return true
	}
}

//roll - rolls the dice and returns the individual dice rolls sum
// Input - consumes rollsCount (how many times dice should be rolled) and diceSize
// fetches each roll value using rand.Intn function
// minimum value of roll = 1 and maximum value of roll = dice size
// Returns - sum of each roll value
func roll(rollsCount int,diceSize int) int{
	var rollSum int
	log.Println("rolling dice size :",diceSize)
	rand.Seed(time.Now().UnixNano())
	for count:=0;count<rollsCount;count++  {
		min := 1
		max := diceSize
		rollValue := rand.Intn(max - min) + min			
		log.Println("roll value: ",rollValue)
		rollSum += rollValue
	}
    return rollSum
}

//RollDices - validates the DiceRoller input string and calls sumRollValue function
//Input : DiceRoller input string (Ex : "4d6+2d10+4d20")
func RollDices(inputString string) {
	ResultChannel := make(chan int)
	if isValidDiceString(inputString)  {
		rolls := splitDices(inputString)
		go sumRollValue(rolls,ResultChannel ) 
	} 
	log.Println("Sum of Roll values : ",<- ResultChannel)
}