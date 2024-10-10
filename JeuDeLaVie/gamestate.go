package main

import (
	rd "math/rand"
)

type GameState struct {
	cell [GAMEY][GAMEY]int
}

func (gs *GameState) RandomInitState() {
	for i := 0; i < GAMEY; i++ {
		for j := 0; j < GAMEX; j++ {
			gs.cell[i][j] = rd.Intn(2)
		}
	}
}

func (gs *GameState) EmptyInitState() {
	for i := 0; i < GAMEY; i++ {
		for j := 0; j < GAMEX; j++ {
			gs.cell[i][j] = 0
		}
	}
}

func (gs *GameState) ComputeNextStep() {
	var buffer [GAMEY][GAMEX]int

	// Counting the neighbours
	for i := 0; i < GAMEY; i++ {
		for j := 0; j < GAMEX; j++ {
			buffer[i][j] = gs.CountNeighbours(i, j)
		}
	}	

	for i := 0; i < GAMEY; i++ {
		for j := 0; j < GAMEX; j++ {
			if (gs.cell[i][j] >= 1) {
				if (buffer[i][j] == 2 || buffer[i][j] == 3) {
					gs.cell[i][j]++
				} else {
					gs.cell[i][j] = 0
				}
			} else {
				if (buffer[i][j] == 3) {
					gs.cell[i][j] = 1
				} else {
					gs.cell[i][j] = 0
				}
			}
		}
	}

}

// Pris de mon projet https://github.com/Yuiko911/gooflife
func (gs GameState) CountNeighbours(py int, px int) int {
	// Based on 
	// https://github.com/Yuiko911/silly/blob/main/game_of_life_clean.c
	
	n := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i == 0 && j == 0) {continue}

			ny := (py + i + GAMEY) % GAMEY;
            nx := (px + j + GAMEX) % GAMEX;

			if gs.cell[ny][nx] >= 1 {n++}
		}
	}

	return n;
}