package main
 
import (
	"html/template"
	"net/http"
	"fmt"
	"os"
	Power4 "Power4/game"
)
 
var templates *template.Template
 
func init() {
	var err error
	templates, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println("Erreur templates :", err)
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/",func(w func(http.ResponseWriter, *http.Request))) {
		var grille [][]string
		Power4.InitGrille(&grille)

		templates.ExecuteTemplate(w, "home",grille)
	}

		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
 
	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


 


