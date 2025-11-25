package main

import (
	Power4 "Power4/game"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"
)

var listTemplates *template.Template
var currentGame Power4.GameState

type Partie struct {
	Joueur1   string
	Joueur2   string
	Vainqueur string
	Date      string
	NbTours   int
}

var historique []Partie
var lastResult Partie

type GamePageData struct {
	Grille       [][]string
	JoueurActuel Power4.Joueur
}

func main() {
	var errTemplate error
	listTemplates, errTemplate = template.ParseGlob("./templates/*.html")
	if errTemplate != nil {
		fmt.Println(errTemplate.Error())
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var grille [][]string
		Power4.InitGrille(&grille)
		err := listTemplates.ExecuteTemplate(w, "base", grille)
		if err != nil {
			http.Error(w, "Erreur template base", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/game/init", func(w http.ResponseWriter, r *http.Request) {
		err := listTemplates.ExecuteTemplate(w, "init", nil)
		if err != nil {
			http.Error(w, "Erreur template init", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/game/init/traitement", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/game/init", http.StatusSeeOther)
			return
		}

		name1 := r.FormValue("player1")
		name2 := r.FormValue("player2")
		color1 := r.FormValue("color1")

		if name1 == "" {
			name1 = "Joueur 1"
		}
		if name2 == "" {
			name2 = "Joueur 2"
		}
		if color1 != "rouge" && color1 != "jaune" {
			color1 = "rouge"
		}

		color2 := "jaune"
		if color1 == "jaune" {
			color2 = "rouge"
		}

		currentGame = Power4.InitGame(name1, name2, color1, color2)
		http.Redirect(w, r, "/game/play", http.StatusSeeOther)
	})

	http.HandleFunc("/game/play", func(w http.ResponseWriter, r *http.Request) {
		if currentGame.Grille == nil {
			http.Redirect(w, r, "/game/init", http.StatusSeeOther)
			return
		}

		data := GamePageData{
			Grille:       currentGame.Grille,
			JoueurActuel: *currentGame.JoueurActuel,
		}

		err := listTemplates.ExecuteTemplate(w, "game", data)
		if err != nil {
			http.Error(w, "Erreur template game", http.StatusInternalServerError)
			return
		}
	})

http.HandleFunc("/game/play/traitement", func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/game/play", http.StatusSeeOther)
		return
	}

	colStr := r.FormValue("col")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Redirect(w, r, "/game/play", http.StatusSeeOther)
		return
	}
	
	col = col - 1

	joueurCourant := currentGame.JoueurActuel

	ok := Power4.JouerCoup(&currentGame, col, joueurCourant.Couleur)
	if !ok {
		http.Redirect(w, r, "/game/play?err=plein", http.StatusSeeOther)
		return
	}

	currentGame.Tour++

	couleurJoueur := joueurCourant.Couleur

	if Power4.CheckWin(currentGame.Grille, couleurJoueur) {

		lastResult = Partie{
			Joueur1:  currentGame.Joueur1.Nom,
			Joueur2:  currentGame.Joueur2.Nom,
			Vainqueur: joueurCourant.Nom,
			Date:     time.Now().Format("02/01/2006 15:04"),
			NbTours:  currentGame.Tour,
		}
		historique = append(historique, lastResult)

		http.Redirect(w, r, "/game/end", http.StatusSeeOther)
		return
	}

	if Power4.GrillePleine(currentGame.Grille) {
		lastResult = Partie{
			Joueur1:  currentGame.Joueur1.Nom,
			Joueur2:  currentGame.Joueur2.Nom,
			Vainqueur: "Égalité",
			Date:     time.Now().Format("02/01/2006 15:04"),
			NbTours:  currentGame.Tour,
		}
		historique = append(historique, lastResult)

		http.Redirect(w, r, "/game/end", http.StatusSeeOther)
		return
	}

	if currentGame.JoueurActuel == &currentGame.Joueur1 {
		currentGame.JoueurActuel = &currentGame.Joueur2
	} else {
		currentGame.JoueurActuel = &currentGame.Joueur1
	}

	http.Redirect(w, r, "/game/play", http.StatusSeeOther)

	})

	http.HandleFunc("/game/end", func(w http.ResponseWriter, r *http.Request) {
		err := listTemplates.ExecuteTemplate(w, "end", lastResult)
		if err != nil {
			http.Error(w, "Erreur template end", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/game/scoreboard", func(w http.ResponseWriter, r *http.Request) {
		err := listTemplates.ExecuteTemplate(w, "scoreboard", historique)
		if err != nil {
			http.Error(w, "Erreur template scoreboard", http.StatusInternalServerError)
			return
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
