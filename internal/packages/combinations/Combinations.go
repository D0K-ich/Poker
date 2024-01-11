package Combinations

import (
	"Poker/internal/models"
	"sort"
)

var fl, str, help = 0, 0, 0

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
			*&pl.WinCombination = 6
			fl++
			return
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
						*&pl.WinCombination = 5
						help = i
						str++
						return
					}
				}
			}
		}
		i++
	}
}

func Fullhouse(pl *models.Player) {
	var para, trip = 0, 0
	if *&pl.WinCombination > 2 {
		return
	}

	//------------------------------------------------------------------------------ StreetFlush
	if str == 1 && fl == 1 {
		if pl.Cards[help].Rank == "10" && pl.Cards[help].Suite == pl.Cards[help+1].Suite {
			if pl.Cards[+1].Rank == "J" && pl.Cards[help].Suite == pl.Cards[help+2].Suite {
				if pl.Cards[help+2].Rank == "Q" && pl.Cards[help].Suite == pl.Cards[help+3].Suite {
					if pl.Cards[help+3].Rank == "K" && pl.Cards[help].Suite == pl.Cards[help+4].Suite {
						if pl.Cards[help+4].Rank == "A" && pl.Cards[help].Suite == pl.Cards[help].Suite {
							*&pl.WinCombination = 10
							return
						}
					}
				}
			}
		}
		*&pl.WinCombination = 9
		return
	}

	//------------------------------------------------------------------------------ Kare
	i := 0
	for i < len(pl.Cards)-3 {
		if pl.Cards[i].Rank == pl.Cards[i+1].Rank && pl.Cards[i+1].Rank == pl.Cards[i+2].Rank && pl.Cards[i+2].Rank == pl.Cards[i+3].Rank {
			*&pl.WinCombination = 8
			return
		}
		i++
	}
	//------------------------------------------------------------------------------ FullHouse
	i = 0
	for i < len(pl.Cards)-2 {
		if pl.Cards[i].Rank == pl.Cards[i+1].Rank && pl.Cards[i+1].Rank == pl.Cards[i+2].Rank {
			trip++
			*&pl.Cards = append(pl.Cards[:i], pl.Cards[i+2:]...)
		}
		i++
	}

	i = 0
	for i < len(pl.Cards)-1 {
		if pl.Cards[i].GetScore(pl.Cards[i]) == pl.Cards[i].GetScore(pl.Cards[i+1]) {
			para += 1
		}
		if trip == 1 && para == 1 {
			*&pl.WinCombination = 7
			return
		}
		i++
	}
	//------------------------------------------------------------------------------Triple
	if trip == 1 && para != 1 {
		*&pl.WinCombination = 4
		return
	}

	i = 0
	//------------------------------------------------------------------------------Para / Rockets
	if para == 1 {
		*&pl.WinCombination = 2
		return
	} else if para == 2 {
		*&pl.WinCombination = 3
		return
	} else if para == 3 {
		*&pl.WinCombination = 2
		return
	}
}

func HighestCard(pl *models.Player) {
	if pl.WinCombination == 0 {
		*&pl.HighestCard = pl.Cards[6]
		*&pl.WinCombination = 1
	}

}
