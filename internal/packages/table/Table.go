package Table

import (
	"Poker/internal/models"
	"fmt"
)

func CreateTables() models.Table {
	Table := models.Table{}
	return Table
}

func ShowPlayers(tab models.Table) {
	for i := 0; i < len(tab.Players); i++ {
		fmt.Print(tab.Players[i].Name)
		for j := 0; j < len(tab.Cards)+2; j++ {
			fmt.Print(tab.Players[i].Cards[j])
		}

		fmt.Println()
	}
}
