package card

import (

	"strconv"
	Pl "Poker/Player"

	"github.com/fatih/color"
)

type Card struct {
	Suite string
	Rank  string
}

var Ranks = getRange(2, 10)
var Suite = SetSuyts()

func MyCard() Card {
	return Card{}
}

func SetSuyts() (suyts [4]string) {
	black := color.New(color.FgMagenta).SprintFunc()
	pick := black("♠")
	pinck := color.New(color.FgRed).SprintFunc()
	heart := pinck("♥")
	yellow := color.New(color.FgYellow).SprintFunc()
	boobs := yellow("♦")
	green := color.New(color.FgGreen).SprintFunc()
	tref := green("♣")
	suyts = [4]string{pick, heart, boobs, tref}

	return suyts
}

func SetRank() (rank []string) {
	rank = getRange(2, 10)
	return rank
}

func getRange(start int, end int) (result []string) {
	result = make([]string, 0)
	for value := start; value <= end; value++ {
		value2 := strconv.Itoa(value)
		result = append(result, value2)
	}
	result = append(result, "J", "Q", "K", "A")
	return result
}

func GetScore(cd Pl.Card) int {
	score := len(Ranks)
	v := cd.Rank
	for i := 0; i < 13; i++ {
		if v == Ranks[i] {
			score = i + 2
		}
	}
	return score
}

func GetSuite(cd Pl.Card) string {
	S := ""
	v := cd.Suite
	for i := 0; i < 4; i++ {
		if v == Suite[i] {
			S = Suite[i]
		}
	}
	return S
}


