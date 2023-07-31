package Deck

import (
	C "Poker/Card"
	"math/rand"
	"time"
)

type Card struct {
	Suite string
	Rank string
}



func Deal(deck1 []Card, deck2 *[]Card) Card {
	lc := deck1[len(deck1)-1]
	DellCard2(deck2)
	return lc
}

func CreateDeck() []Card {
	suyts := C.SetSuyts()
	rank := C.SetRank()
	Deck := []Card{}
	for i := range suyts {
		for j := range rank {
			card := Card{suyts[i], rank[j]}
			Deck = append(Deck, card)
		}
	}

	ShuffleDeck(Deck)
	return Deck
}

func ShuffleDeck(Deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Deck), func(i, j int) { Deck[i], Deck[j] = Deck[j], Deck[i] })
}

func DellCard2(deck *[]Card) {
	*deck = (*deck)[:len(*deck)-1]
}
