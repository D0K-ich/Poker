package game

import (
	Com "Poker/Combinations"
	Deck "Poker/Deck"
	Pl "Poker/Player"
	Tb "Poker/Table"
	"fmt"
)

type Card struct {
	Suite string
	Rank  string
}

var MainDeck = Deck.CreateDeck()

func JoinPlayer(player *Pl.Player, table *Tb.Table) {
	*&table.Players = append(*&table.Players, *player)
}

func Startgame(table *Tb.Table, pl *Pl.Player) {

	for i := 0; i < 1; i++ {
		Pl.GiveCard(&*pl, Pl.Card(Deck.Deal(MainDeck, &MainDeck)))
		Pl.GiveCard(&*pl, Pl.Card(Deck.Deal(MainDeck, &MainDeck)))
	}

	Flop(table)
	Turn(table)
	River(*&table, *&pl)

}

func Endgame(pl *Pl.Player, t Tb.Table) {

	Com.Street(*&pl)
	Com.Flash(*&pl, Com.GetSuitMap(t))
	Com.Fullhouse(*&pl)
	//*&pl.WinCombination = 666

	fmt.Print(*&pl.WinCombination)
}

func Flop(table *Tb.Table) {
	table.Cards = append(table.Cards, Tb.Card(Deck.Deal(MainDeck, &MainDeck)))
	table.Cards = append(table.Cards, Tb.Card(Deck.Deal(MainDeck, &MainDeck)))
	table.Cards = append(table.Cards, Tb.Card(Deck.Deal(MainDeck, &MainDeck)))
	//fmt.Println("Стол: ", table.Cards[0], table.Cards[1], table.Cards[2])
}

func Turn(table *Tb.Table) {
	table.Cards = append(table.Cards, Tb.Card(Deck.Deal(MainDeck, &MainDeck)))
	//fmt.Println("Стол: ", table.Cards[0], table.Cards[1], table.Cards[2], table.Cards[3])
}

func River(table *Tb.Table, p *Pl.Player) {
	table.Cards = append(table.Cards, Tb.Card(Deck.Deal(MainDeck, &MainDeck)))
	fmt.Println("Стол: ", table.Cards[0], table.Cards[1], table.Cards[2], table.Cards[3], table.Cards[4])

	for i := 0; i < 5; i++ {
		Pl.GiveCard(&*p, Pl.Card(table.Cards[i]))
	}

	*&table.Players[0].Cards = p.Cards
}
