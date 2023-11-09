package Player

//Deck "Poker/deck"

//"fmt"

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

func GiveCard(pl1 *Player, Cards Card) {
	*&pl1.Cards = append(*&pl1.Cards, Cards)
}

func AddNewPlayer(cash int, name string) Player {
	PlayerNew := Player{}
	PlayerNew.Name = name
	PlayerNew.Cash = cash
	PlayerNew.Cards = []Card{}
	return PlayerNew
}