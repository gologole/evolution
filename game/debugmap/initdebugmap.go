package debugmap

import "main.go/game/initmap"

var id = 0

func InitDebugMap(newmap initmap.GlobalMap, energy int, ilumitium int, organic int, mutantchance int) {
	newmap.Height = 100
	newmap.Width = 100
	newmap.Energy = energy
	newmap.Ilumitium = ilumitium
	newmap.Organic = organic
	newmap.Mutantchance = mutantchance
	newmap.Id = id + 1
	InitandFillChunk(newmap) //заполняет чанки дефолтными клетками

}

func InitandFillChunk(newmap initmap.GlobalMap) {
	newmap.Chanks = make([][]initmap.Chank, 2)
	for i := 0; i < 2; i++ {
		newmap.Chanks[i] = make([]initmap.Chank, 2)
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			newmap.Chanks[i][j] = initmap.Chank{
				i + j,
				fillChank(newmap),
			}
		}
	}
}

func fillChank(newmap initmap.GlobalMap) [][]initmap.Cell {
	//init cell map in chunk
	cells := make([][]initmap.Cell, 25)
	for i := 0; i < 25; i++ {
		cells[i] = make([]initmap.Cell, 25)
	}

	//fill cells
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			cell := initmap.InitCell(newmap, i, j)
			cells[i][j] = *cell
		}
	}
	return cells
}
