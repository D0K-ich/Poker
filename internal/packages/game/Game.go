package game

import (
	"Poker/internal/models"
	Com "Poker/internal/packages/Combinations"
	"fmt"
)

var MainDeck = models.Deck{}.CreateDeck()

func Startgame(table *models.Table, pl *models.Player) {
	for i := 0; i < 1; i++ {
		models.Card{}.GiveCard(&*pl, models.Card(models.Deck{}.Deal(MainDeck, &MainDeck)))
		models.Card{}.GiveCard(&*pl, models.Card(models.Deck{}.Deal(MainDeck, &MainDeck)))
	}
	Flop(table)
	Turn(table)
	River(*&table, *&pl)
}

func Endgame(pl *models.Player, t models.Table) {
	Com.Street(*&pl)
	Com.Flash(*&pl, Com.GetSuitMap(t))
	Com.Fullhouse(*&pl)
	//*&pl.WinCombination = 666
	fmt.Print(*&pl.WinCombination)
}

func Flop(table *models.Table) {
	table.Cards = append(table.Cards, models.Card(models.Deck{}.Deal(MainDeck, &MainDeck)))
	table.Cards = append(table.Cards, models.Card(models.Deck{}.Deal(MainDeck, &MainDeck)))
	table.Cards = append(table.Cards, models.Card(models.Deck{}.Deal(MainDeck, &MainDeck)))
	//fmt.Println("Стол: ", table.Cards[0], table.Cards[1], table.Cards[2])
}

func Turn(table *models.Table) {
	table.Cards = append(table.Cards, models.Card(models.Deck{}.Deal(MainDeck, &MainDeck)))
	//fmt.Println("Стол: ", table.Cards[0], table.Cards[1], table.Cards[2], table.Cards[3])
}

func River(table *models.Table, p *models.Player) {
	table.Cards = append(table.Cards, models.Card(models.Deck{}.Deal(MainDeck, &MainDeck)))
	fmt.Println("Стол: ", table.Cards[0], table.Cards[1], table.Cards[2], table.Cards[3], table.Cards[4])
	for i := 0; i < 5; i++ {
		models.Card{}.GiveCard(&*p, models.Card(table.Cards[i]))
	}
	*&table.Players[0].Cards = p.Cards
}
