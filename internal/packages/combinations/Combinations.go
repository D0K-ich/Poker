package Combinations

import (
	"Poker/internal/models"

	"sort"
)

var DD, trip = 0, 0

func SortCards(pl *models.Player) {
	i := 0
	j := 0
	for i < len(pl.Cards)-1 {
		for j < len(pl.Cards)-1 {
			if (models.Card{}.GetScore(*&pl.Cards[j]) > models.Card{}.GetScore(*&pl.Cards[j+1])) {
				*&pl.Cards[j], *&pl.Cards[j+1] = pl.Cards[j+1], pl.Cards[j]
			}
			j++
		}
		j = 0
		i++
	}
}

func GetSuitMap(t models.Table) map[string]int {
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

func Flash(pl *models.Player, cm map[string]int) {
	i := 0
	for range cm {

		if cm[pl.Cards[i].Suite] == 5 {
			*&pl.WinCombination = 5
		}
		i++
	}
}

func Street(pl *models.Player) {
	i := 0
	for i < len(pl.Cards)-4 {
		if (models.Card{}.GetScore(pl.Cards[i])+1 == models.Card{}.GetScore(pl.Cards[i+1])) {
			if (models.Card{}.GetScore(pl.Cards[i+1])+1 == models.Card{}.GetScore(pl.Cards[i+2])) {
				if (models.Card{}.GetScore(pl.Cards[i+2])+1 == models.Card{}.GetScore(pl.Cards[i+3])) {
					if (models.Card{}.GetScore(pl.Cards[i+3])+1 == models.Card{}.GetScore(pl.Cards[i+4])) {
						*&pl.WinCombination = 4
					}
				}
			}
		}
		i++
	}
}

func Fullhouse(pl *models.Player) {

	i := 0
	ch := 0
	for i < len(pl.Cards)-1 {
		if pl.Cards[i].Rank == pl.Cards[i+1].Rank {
			ch++
			DD++
		}
		i++
	}
	if ch == 1 {
		*&pl.WinCombination = 1
	} else if ch == 2 {
		*&pl.WinCombination = 2
	}

	i = 0
	for i < len(pl.Cards)-2 {
		if pl.Cards[i].Rank == pl.Cards[i+1].Rank && pl.Cards[i+1].Rank == pl.Cards[i+2].Rank {
			*&pl.WinCombination = 3
			trip++
		}
		i++
	}

	if DD > 0 && trip > 0 {
		*&pl.WinCombination = 6
	}
}
