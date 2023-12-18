package utilites

import (
	"github.com/fatih/color"
	"strconv"
)

func GetRange(start int, end int) (result []string) {
	result = make([]string, 0)
	for value := start; value <= end; value++ {
		value2 := strconv.Itoa(value)
		result = append(result, value2)
	}
	result = append(result, "J", "Q", "K", "A")
	return result
}

func SetRank() (rank []string) {
	rank = GetRange(2, 10)
	return rank
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
