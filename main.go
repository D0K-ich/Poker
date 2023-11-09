package main

import (
	Com "Poker/Combinations"
	"Poker/Deck"
	game "Poker/Game"
	Pl "Poker/Player"
	Tab "Poker/Table"
)

var MainDeck = Deck.CreateDeck()

func main() {
	i := 0
	for i == 0 {
		table1 := Tab.CreateTables()

		Danches := Pl.AddNewPlayer(1000, "Danches")

		game.JoinPlayer(&Danches, &table1)
		game.Startgame(&table1, &Danches)

		Com.SortCards(&Danches)

		Tab.ShowPlayers(table1)

		game.Endgame(&Danches, table1)

		if Danches.WinCombination == 6 {
			i = 1
		}
	}

}
