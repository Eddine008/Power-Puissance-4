package Power4

type Joueur struct {
	Nom     string
	Couleur string
	Signe   string
}

type GameState struct {
	Grille       [][]string
	Joueur1      Joueur
	Joueur2      Joueur
	JoueurActuel *Joueur
	Tour         int
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

func InitGame(name1, name2, color1, color2 string) GameState {
	var game GameState

	InitGrille(&game.Grille)

	game.Joueur1.Nom = name1
	game.Joueur1.Couleur = color1
	game.Joueur1.Signe = "O"

	game.Joueur2.Nom = name2
	game.Joueur2.Couleur = color2
	game.Joueur2.Signe = "X"

	game.JoueurActuel = &game.Joueur1
	game.Tour = 0

	return game
}

func JouerCoup(game *GameState, col int, couleur string) bool {
	if col < 0 || col > 7 {
		return false
	}

	for ligne := len(game.Grille) - 1; ligne >= 0; ligne-- {
		if game.Grille[ligne][col] == "-" {
			game.Grille[ligne][col] = couleur
			return true
		}
	}

	return false

}

func GrillePleine(grille [][]string) bool {
	for i := 0; i < len(grille); i++ {
		for j := 0; j < len(grille[i]); j++ {
			if grille[i][j] == "-" {
				return false
			}
		}
	}
	return true
}

func CheckWin(grille [][]string, couleur string) bool {
	rows := len(grille)
	cols := len(grille[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols-3; j++ {
			if grille[i][j] == couleur &&
				grille[i][j+1] == couleur &&
				grille[i][j+2] == couleur &&
				grille[i][j+3] == couleur {
				return true
			}
		}
	}

	for j := 0; j < cols; j++ {
		for i := 0; i < rows-3; i++ {
			if grille[i][j] == couleur &&
				grille[i+1][j] == couleur &&
				grille[i+2][j] == couleur &&
				grille[i+3][j] == couleur {
				return true
			}
		}
	}

	for i := 0; i < rows-3; i++ {
		for j := 0; j < cols-3; j++ {
			if grille[i][j] == couleur &&
				grille[i+1][j+1] == couleur &&
				grille[i+2][j+2] == couleur &&
				grille[i+3][j+3] == couleur {
				return true
			}
		}
	}

	for i := 3; i < rows; i++ {
		for j := 0; j < cols-3; j++ {
			if grille[i][j] == couleur &&
				grille[i-1][j+1] == couleur &&
				grille[i-2][j+2] == couleur &&
				grille[i-3][j+3] == couleur {
				return true
			}
		}
	}

	return false
}
