package initmap

import "C"
import (
	"main.go/game/genom"
	"sync"
	"unsafe"
)

const (
	deafultgenomheinght = 10
	deafultgenomwidth   = 10
	deafultmaxnumgenom  = 100
	deafulthp           = 10
	deafultmaxage       = 30
	deafultneedenergy   = 10
	deafultmaxenergy    = 100
	deafultenergynow    = 100
)

type Chank struct {
	Chankid int
	Cells   [][]Cell
}

type GlobalMap struct {
	//Gamesparametrs
	Height              int
	Width               int
	Id                  int
	Chanks              [][]Chank
	DistanceBtwnewCells int
	Organic             int
	Energy              int
	Ilumitium           int
	Mutantchance        int //1-100 вероятность мутации
	Cells               [][]Cell
}

type Cell struct {
	IsLife  int
	X_coord int
	Y_coord int

	//not life
	Organic   int
	Energy    int
	Ilumitium int

	//life
	OrganizmID            int
	Activgen              int
	Genom                 [][]int
	Typeofcell            int
	Hp                    int
	Age                   int
	Maxage                int
	Needenergy            int //потребление в один тик
	Maxenergy             int //максимальный запас
	Energynow             int
	OrganicAfterdie       int
	EnergyAfterdie        int
	EnengytransportParams energytransport
}

type energytransport struct {
	x_parentscoords int
	y_parentscoords int
}

func GlobalMapconstructor(id int, height int, width int, distanceBtwnewCells int, organic int, energy int, ilumitium int, mutantchance int) *GlobalMap {
	var Map GlobalMap
	Map.Id = id
	Map.Height = height
	Map.Width = width
	Map.DistanceBtwnewCells = distanceBtwnewCells + 1
	Map.Organic = organic
	Map.Energy = energy
	Map.Ilumitium = ilumitium
	Map.Mutantchance = mutantchance
	return &Map
}

func InitCell(newmap GlobalMap, x int, y int) *Cell {
	var MyCell Cell

	MyCell.IsLife = 0
	MyCell.Organic = newmap.Organic
	MyCell.Energy = newmap.Energy
	MyCell.Ilumitium = newmap.Ilumitium
	MyCell.X_coord = x
	MyCell.Y_coord = y
	return &MyCell
}

func InitcellsMap(newmap GlobalMap) [][]Cell {
	//Cells := make([][]Cell, newmap.Height)
	Cells := make([][]Cell, newmap.Height)
	for i := range Cells {
		Cells[i] = make([]Cell, newmap.Width)
	}
	distance := newmap.DistanceBtwnewCells

	wg := sync.WaitGroup{}
	wg.Add(newmap.Height * newmap.Width)

	for x := 0; x < newmap.Height; x++ {
		for y := 0; y < newmap.Width; y++ {
			func(x int, y int, distance int, cells [][]Cell) {
				newcell := InitCell(newmap, x, y)
				if x%distance == 0 && y%distance == 0 { //клетка живая
					newcell.IsLife = 1
					newcell.Age = 0
					newcell.Hp = deafulthp
					newcell.Maxage = deafultmaxage
					newcell.Needenergy = deafultneedenergy
					newcell.Maxenergy = deafultmaxenergy
					newcell.Energynow = deafultenergynow
					newcell.Typeofcell = 1

					newgenom := make([][]int, newmap.Height)
					for i := range Cells {
						newgenom[i] = make([]int, newmap.Width)
					}
					c_point := genom.InitGenom(newgenom, deafultgenomheinght, int(deafultgenomwidth), int(deafultmaxnumgenom))
					newcell.Genom = newgenom
					genom.FreeMemoryByC(unsafe.Pointer(c_point), deafultgenomheinght)
					Cells[x][y] = *newcell
					wg.Done()
				} else { //клетка не живая
					Cells[x][y] = *newcell
					wg.Done()
				}
			}(x, y, distance, Cells)
		}
	}

	wg.Wait()
	return Cells
}
