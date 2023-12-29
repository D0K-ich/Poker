package models

import (
	"Poker/internal/packages/utilites"
	"math/rand"
	"time"
)

type Card struct {
	Suite string
	Rank  string
}
type Player struct {
	Cash           int
	Name           string
	Cards          []Card
	WinCombination int
}
type Table struct {
	Cards   []Card
	Players []Player
}
type Deck struct {
	Cards []Card
}

//------------------------------------------------------------------------------

var Ranks = utilites.GetRange(2, 10)
var Suite = utilites.SetSuyts()

//------------------------------------------------------------------------------

func (Player) JoinPlayer(player *Player, table *Table) {
	*&table.Players = append(*&table.Players, *player)
}
func (Player) AddNewPlayer(cash int, name string) Player {
	PlayerNew := Player{}
	PlayerNew.Name = name
	PlayerNew.Cash = cash
	PlayerNew.Cards = []Card{}
	return PlayerNew
}

//------------------------------------------------------------------------------

func (Card) GiveCard(pl1 *Player, Cards Card) {
	*&pl1.Cards = append(*&pl1.Cards, Cards)
}
func (Card) GetScore(cd Card) int {
	score := len(Ranks)
	v := cd.Rank
	for i := 0; i < 13; i++ {
		if v == Ranks[i] {
			score = i + 2
		}
	}
	return score
}
func (Card) GetSuite(cd Card) string {
	S := ""
	v := cd.Suite
	for i := 0; i < 4; i++ {
		if v == Suite[i] {
			S = Suite[i]
		}
	}
	return S
}

//------------------------------------------------------------------------------

func (Deck) ShuffleDeck(Deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Deck), func(i, j int) { Deck[i], Deck[j] = Deck[j], Deck[i] })
}

func (Deck) DellCard2(deck *[]Card) {
	*deck = (*deck)[:len(*deck)-1]
}

func (Deck) Deal(deck1 []Card, deck2 *[]Card) Card {
	lc := deck1[len(deck1)-1]
	Deck{}.DellCard2(deck2)
	return lc
}

func (Deck) CreateDeck() []Card {
	suyts := utilites.SetSuyts()
	rank := utilites.SetRank()
	deck := []Card{}
	for i := range suyts {
		for j := range rank {
			card := Card{suyts[i], rank[j]}
			deck = append(deck, card)
		}
	}

	Deck{}.ShuffleDeck(deck)
	return deck
}
