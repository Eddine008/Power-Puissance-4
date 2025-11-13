package Power4

type Joueur struct {
	Nom     string
	Couleur string
	Signe   string
}

func InitGrille(grille *[][]string) {
	rows, cols := 6, 7
	*grille = make([][]string, rows)
	for i := 0; i < rows; i++ {
		(*grille)[i] = make([]string, cols)
		for j := 0; j < cols; j++ {
			(*grille)[i][j] = "-"
		}
	}
}

func InitPlayers(player1 *Joueur, player2 *Joueur) {
	player1.Nom = "Joueur 1"
	player1.Couleur = "Rouge"
	player1.Signe = "O"

	player2.Nom = "Joueur 2"
	player2.Couleur = "Jaune"
	player2.Signe = "X"
}
