package Combinations

import (
	Pl "Poker/Player"
	Tb "Poker/Table"
	"fmt"
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

func Flash(pl *Pl.Player, cm map[string]int) {
	i := 0
	for range cm {

		if cm[pl.Cards[i].Suite] == 5 {
			*&pl.WinCombination = 1
		}
		i++
	}

}

func Double(pl *Pl.Player) {
	i := 0
	for range pl.Cards {
		if pl.Cards[i].Rank == pl.Cards[i+1].Rank {
			fmt.Print("Пара")
		}
		i++
	}
}

func SortCards(pl *Pl.Player) {
	i := 0
	for range pl.Cards {
		if pl.Cards[i].Rank > pl.Cards[i+1].Rank {
			*&pl.Cards[i] = pl.Cards[i+1]
		}
		i++
	}
}
