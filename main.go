package main

import (
	"fmt"
	//Cd "Poker/Card"
	Com "Poker/Combinations"
	"Poker/Deck"
	game "Poker/Game"
	Pl "Poker/Player"
	Tab "Poker/Table"
)

var MainDeck = Deck.CreateDeck()

func main() {
	table1 := Tab.CreateTables()

	Danches := Pl.AddNewPlayer(1000, "Danches")
	game.JoinPlayer(&Danches, &table1)
	game.Startgame(&table1, &Danches)

	Tab.ShowPlayers(table1)

	fmt.Println(Com.GetSuitMap(table1))

	//Cd.CountEqCard(Danches.Cards[1], "♠")
}
