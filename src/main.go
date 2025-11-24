package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var listTemplates *template.Template

func main() {

	var errTemplate error
	listTemplates, errTemplate = template.ParseGlob("./templates/*.html")
	if errTemplate != nil {
		fmt.Println(errTemplate.Error())
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// var grille [][]string
		// Power4.InitGrille(&grille)
		//err := listTemplates.ExecuteTemplate(w, "base", grille)
		// if err != nil {
		// 	http.Error(w, "Erreur template base", http.StatusInternalServerError)
		// 	fmt.Println("Erreur template base :", err)
		// 	return
		// }
		listTemplates.ExecuteTemplate(w, "base", nil)
	})

	http.HandleFunc("/game/play", func(w http.ResponseWriter, r *http.Request) {
		// var grille [][]string
		// Power4.InitGrille(&grille)
		err := listTemplates.ExecuteTemplate(w, "game", grille)
		if err != nil {
			http.Error(w, "Erreur template game", http.StatusInternalServerError)
			fmt.Println("Erreur template game :", err)
			return
		}
	})

	http.HandleFunc("/game/scoreboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Page scoreboard (à compléter).")
	})

	http.HandleFunc("/game/init", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Page d'initialisation (à compléter avec un formulaire).")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
