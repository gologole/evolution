package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"main.go/game/debugmap"
	"main.go/game/initmap"
	"net/http"
	"strconv"
)

var newmap = initmap.GlobalMap{}
var url string

func RunDebug() {
	fmt.Println("set port")
	var port int
	fmt.Scan(&port)
	fmt.Println("set frontend url")
	fmt.Scan(&url)

	http.HandleFunc("/debug/initmap", initmapHandler)
	http.HandleFunc("/debug/setCell", setCellHandler)
	http.HandleFunc("/debugsite", staticHandler)

	http.ListenAndServe(":"+strconv.Itoa(port), nil)

	fmt.Println("Debug server is start")
}

func setCellHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cell := initmap.Cell{}

		cell.OrganizmID, _ = strconv.Atoi(r.FormValue("OrganizmID"))
		cell.Activgen, _ = strconv.Atoi(r.FormValue("Activgen"))
		cell.Typeofcell, _ = strconv.Atoi(r.FormValue("Typeofcell"))
		cell.Hp, _ = strconv.Atoi(r.FormValue("Hp"))
		cell.Age, _ = strconv.Atoi(r.FormValue("Age"))
		cell.Maxage, _ = strconv.Atoi(r.FormValue("Maxage"))
		cell.X_coord, _ = strconv.Atoi(r.FormValue("X_coord"))
		cell.Y_coord, _ = strconv.Atoi(r.FormValue("Y_coord"))
		cell.Needenergy, _ = strconv.Atoi(r.FormValue("Needenergy"))
		cell.Maxenergy, _ = strconv.Atoi(r.FormValue("Maxenergy"))
		cell.Energynow, _ = strconv.Atoi(r.FormValue("Energynow"))
		cell.OrganicAfterdie, _ = strconv.Atoi(r.FormValue("OrganicAfterdie"))
		cell.EnergyAfterdie, _ = strconv.Atoi(r.FormValue("EnergyAfterdie"))

		//обновление поля
		debugmap.SetCell(newmap, cell.X_coord, cell.Y_coord, cell)

		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				jsonData, err := json.Marshal(newmap.Chanks[i][j])
				if err != nil {
					log.Fatalf("Failed to encode data to JSON: %v", err)
				}

				resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
				if err != nil {
					log.Fatalf("Failed to make POST request: %v", err)
				}
				defer resp.Body.Close()

				log.Println("POST request sent successfully")
			}
		}
	}
}

func initmapHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		energy, _ := strconv.Atoi(r.FormValue("energy"))
		ilumitium, _ := strconv.Atoi(r.FormValue("ilumitium"))
		organic, _ := strconv.Atoi(r.FormValue("organic"))
		mutantchance, _ := strconv.Atoi(r.FormValue("mutantchance"))
		debugmap.InitDebugMap(newmap, energy, ilumitium, organic, mutantchance)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "jsjs.html")
}
