package Power4

type Grille struct {
	Board [][]string
}

func addPion(grille *Grille, col int, player Player) bool {

	nLigne := -1

	for i := len(grille.Board) - 1; i > 0; i-- {
		if grille.Board[i][col-1] == " " {
			nLigne = i
			break
		}
	}

	if nLigne == -1 {
		return false
	}

	grille.Board[nLigne][col-1] = player.Color
	return true
}
