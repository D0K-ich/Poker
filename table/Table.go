package tables

import ( Pl "Poker/Player"
		"fmt"
)

type Card struct {
	Suite string
	Rank string
}

type Table struct {
	Cards []Card
	Players []Pl.Player
}

func CreateTables() Table {
	Table := Table{}
	return Table
}

func ShowPlayers(tab Table) {
	for i := 0; i < len(tab.Players) ; i++ {
		fmt.Print(tab.Players[i].Name)
		for j := 0; j < len(tab.Cards)+2; j++ {
			fmt.Print(tab.Players[i].Cards[j])		
		}

		fmt.Println()
	}
}



