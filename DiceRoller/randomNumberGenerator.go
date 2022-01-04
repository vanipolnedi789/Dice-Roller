package DiceRoller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type internalRandNumGeneratorSource struct{}

type httpRandNumGeneratorSource struct {
	url      string
	response interface{}
}

type randomGenerator interface {
	roll(int, int) int
}

//roll - rolls the dice and returns the individual dice rolls sum
// Input - consumes rollsCount (how many times dice should be rolled) and diceSize
// fetches each roll value using rand.Intn function
// minimum value of roll = 1 and maximum value of roll = dice size
// Returns - sum of each roll value
func (random internalRandNumGeneratorSource) roll(rollsCount int, diceSize int) int {
	var rollSum int
	log.Println("rolling dice size :", diceSize)
	rand.Seed(time.Now().UnixNano())
	for count := 0; count < rollsCount; count++ {
		min := 1
		max := diceSize
		rollValue := rand.Intn(max-min) + min
		log.Println("roll value: ", rollValue)
		rollSum += rollValue
	}
	return rollSum
}

//roll - rolls the dice and returns the individual dice rolls sum
// Input - consumes rollsCount (how many times dice should be rolled) and diceSize
// fetches each roll value using random number generator api
// minimum value of roll = 1 and maximum value of roll = dice size
// Returns - sum of each roll value
func (httpSource httpRandNumGeneratorSource) roll(rollsCount int, diceSize int) int {
	log.Println("rolling dice size :", diceSize)
	httpSource.url = "http://www.randomnumberapi.com/api/v1.0/random?min=1&max=" + fmt.Sprint(diceSize) + "&count=" + fmt.Sprint(rollsCount)

	resp, err := http.Get(httpSource.url)
	if err != nil {
		log.Fatalln("Unable to send request to random number generator api : ", err)
	}

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read random number generator api response: ", err)
	}

	//Convert the body to type int
	unmarshalErr := json.Unmarshal(body, &httpSource.response)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	var rollSum int
	for _, k := range httpSource.response.([]interface{}) {
		log.Println("roll value: ", k)
		rollSum += int(k.(float64))
	}
	return rollSum
}

//rollIndividualDice - calls the roll function based on the random number generator type received
func rollIndividualDice(randNumGenerator randomGenerator, rollsCount int, diceSize int) int {
	return randNumGenerator.roll(rollsCount, diceSize)
}