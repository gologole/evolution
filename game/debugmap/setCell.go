package debugmap

import "main.go/game/initmap"

func SetCell(newmap initmap.GlobalMap, x int, y int, cell initmap.Cell) {
	var chankx int
	var chanky int
	if x < 24 {
		chankx = 0
	}
	if x > 23 {
		chankx = 1
	}
	if y < 24 {
		chanky = 0
	}
	if y > 23 {
		chanky = 1
	}
	newmap.Chanks[chankx][chanky].Cells[x%25][y%25] = cell

}
