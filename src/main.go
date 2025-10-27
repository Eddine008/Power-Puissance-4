package main

import (
	Power4 "Power4/game"
	"fmt"
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
	var maGrille [][]string
	Power4.InitGrille(&maGrille)
	for _, ligne := range maGrille {
		fmt.Println(ligne)
	}
}
