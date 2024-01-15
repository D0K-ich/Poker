package main

import (
	"Poker/internal/models"
	"Poker/internal/packages/combinations"
	"Poker/internal/packages/game"
	Tab "Poker/internal/packages/table"
)

var MainDeck = models.Deck{}.CreateDeck()

func main() {

	Test()
}

func Test() {
	i := 0

	for i == 0 {
		table1 := Tab.CreateTables()

		Danches := models.Table{}.AddNewPlayer(1000, "Danches")
		//Vanches := models.Table{}.AddNewPlayer(1000, "Vanches")

		models.Table{}.JoinPlayer(&Danches, &table1)
		//models.Table{}.JoinPlayer(&Vanches, &table1)

		game.Startgame(&table1, &Danches)

		combinations.SortCards(&Danches)

		Tab.ShowPlayers(table1)

		game.Endgame(&Danches, table1)

		models.Deck{}.Refresh(&MainDeck)

		i++

		if Danches.WinCombination == 1 {
			return
		}
	}
}
