# Dice-Roller
Rolls the different dices and returns the sum of all dices roll values

Go packages used to solve this problem:

-	"math/rand"
-	"time"
-	"strings"
-	"strconv"
-  "log"


To Run the project :
 - go run main.go

To change the random number generation source : 
 - currently package is using internal source for random number generation.
 - To use http api for random number generation , change **randomGenerator** variable datatype to **httpRandNumGeneratorSource**
