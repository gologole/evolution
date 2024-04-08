package game

import (
	"fmt"
	"main.go/game/initmap"
)

func TestInitMap() {
	globalmap := initmap.GlobalMapconstructor(1, 100, 100, 5, 50, 50, 10, 40)
	initcells := initmap.InitcellsMap(*globalmap)
	for i := 0; i < len(initcells); i++ {
		for j := 0; j < len(initcells[i]); j++ {
			fmt.Println(initcells[i][j].IsLife)
		}
	}
}
