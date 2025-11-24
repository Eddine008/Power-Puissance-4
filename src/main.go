package main

import (
	Power4 "Power4/game"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// func initGrille(grille *[][]string) {
// 	rows, cols := 6, 7
// 	*grille = make([][]string, rows)
// 	for i := 0; i < rows; i++ {
// 		(*grille)[i] = make([]string, cols)
// 		for j := 0; j < cols; j++ {
// 			(*grille)[i][j] = "-"
// 		}
// 	}
// }

func main() {
	// var maGrille [][]string
	// Power4.InitGrille(&maGrille)
	// for _, ligne := range maGrille {
	// 	fmt.Println(ligne)
	// test changement

	listTemplates, errTemplate := template.ParseGlob("./templates/*.html")
	if errTemplate != nil {
		fmt.Println(errTemplate.Error())
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var grille [][]string
		Power4.InitGrille(&grille)

		listTemplates.ExecuteTemplate(w, "base", nil)
	})
	//INITIALISATION DU SERVEUR//
	http.ListenAndServe("localhost:8080", nil)
}
