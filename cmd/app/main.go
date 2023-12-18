package main

import (
	"Poker/internal/models"
	Com "Poker/internal/packages/Combinations"
	game "Poker/internal/packages/Game"
	Tab "Poker/internal/packages/Table"
)

var MainDeck = models.Deck{}.CreateDeck()

func main() {
	i := 0
	for i == 0 {
		table1 := Tab.CreateTables()

		Danches := models.Player{}.AddNewPlayer(1000, "Danches")

		models.Player{}.JoinPlayer(&Danches, &table1)
		game.Startgame(&table1, &Danches)

		Com.SortCards(&Danches)

		Tab.ShowPlayers(table1)

		game.Endgame(&Danches, table1)

		if Danches.WinCombination == 6 {
			i = 1
		}
	}

}
