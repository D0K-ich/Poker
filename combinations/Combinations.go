package combinations

import (
	"fmt"
	Tb "Poker/Table"
	"sort"
)


func GetSuitMap(t Tb.Table) map[string]int {
	suitmass := make([]string, 0)

	for i := 0; i < len(t.Players[0].Cards); i++ {
		suitmass = append(suitmass, t.Players[0].Cards[i].Suite)
	}

	usuitmass := removeDuplicates(suitmass)

	sort.Strings(suitmass)

	cardMass := map[string]int{}

	for v := 0; v < len(usuitmass); v++ {
		cardMass[usuitmass[v]] = CuclSuit(suitmass, usuitmass[v])
	}

	fmt.Println(cardMass)

	return cardMass
}

func CuclSuit(arr []string, s string) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == s {
			result++
		}
	}
	return result
}

func removeDuplicates(arr []string) []string {
	result := make([]string, 0, len(arr))
	seen := make(map[string]struct{})
	for _, val := range arr {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = struct{}{}
		}
	}
	return result
}


