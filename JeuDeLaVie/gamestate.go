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

func (gs *GameState) PutCellLoop(y int, x int) {
	ny := (y + GAMEY) % GAMEY
	nx := (x + GAMEX) % GAMEX

	gs.cell[ny][nx] = 1;
}

func (gs *GameState) ToggleCellAtCursor(gv GameView) {
	gs.cell[gv.cy][gv.cx] = (gs.cell[gv.cy][gv.cx] + 1) % 2 
}

func (gs *GameState) AddPatternAtCursor(gv GameView) {
	// oskour
	switch(gv.pat) {
	// Single cell 
	case 0:
		gs.ToggleCellAtCursor(gv)
	// Still lifes
	case 1: // Block
		gs.PutCellLoop(gv.cy+0, gv.cx+0)
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+1)
	case 2: // Beehive
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+0, gv.cx+2)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+3)
		gs.PutCellLoop(gv.cy+2, gv.cx+1)
		gs.PutCellLoop(gv.cy+2, gv.cx+2)
	case 3: // Loaf
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+0, gv.cx+2)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+3)
		gs.PutCellLoop(gv.cy+2, gv.cx+1)
		gs.PutCellLoop(gv.cy+2, gv.cx+3)
		gs.PutCellLoop(gv.cy+3, gv.cx+2)
	case 4: // Boat
		gs.PutCellLoop(gv.cy+0, gv.cx+0)
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+2)
		gs.PutCellLoop(gv.cy+2, gv.cx+1)
	case 5: // Tub
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+2)
		gs.PutCellLoop(gv.cy+2, gv.cx+1)
	// Oscillators
	case 6: // Blinker
		gs.PutCellLoop(gv.cy+0, gv.cx+0)
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+0, gv.cx+2)
	case 7: // Toad
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+0, gv.cx+2)
		gs.PutCellLoop(gv.cy+0, gv.cx+3)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+1)
		gs.PutCellLoop(gv.cy+1, gv.cx+2)
	case 8: // Beacon
		gs.PutCellLoop(gv.cy+0, gv.cx+0)
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+1)
		gs.PutCellLoop(gv.cy+2, gv.cx+2)
		gs.PutCellLoop(gv.cy+2, gv.cx+3)
		gs.PutCellLoop(gv.cy+3, gv.cx+2)
		gs.PutCellLoop(gv.cy+3, gv.cx+3)
	case 9: // Pulsat
		gs.PutCellLoop(gv.cy+0, gv.cx+2)
		gs.PutCellLoop(gv.cy+0, gv.cx+3)
		gs.PutCellLoop(gv.cy+0, gv.cx+4)
		gs.PutCellLoop(gv.cy+0, gv.cx+8)
		gs.PutCellLoop(gv.cy+0, gv.cx+9)
		gs.PutCellLoop(gv.cy+0, gv.cx+10)
		gs.PutCellLoop(gv.cy+2, gv.cx+0)
		gs.PutCellLoop(gv.cy+2, gv.cx+5)
		gs.PutCellLoop(gv.cy+2, gv.cx+7)
		gs.PutCellLoop(gv.cy+2, gv.cx+12)
		gs.PutCellLoop(gv.cy+3, gv.cx+0)
		gs.PutCellLoop(gv.cy+3, gv.cx+5)
		gs.PutCellLoop(gv.cy+3, gv.cx+7)
		gs.PutCellLoop(gv.cy+3, gv.cx+12)
		gs.PutCellLoop(gv.cy+4, gv.cx+0)
		gs.PutCellLoop(gv.cy+4, gv.cx+5)
		gs.PutCellLoop(gv.cy+4, gv.cx+7)
		gs.PutCellLoop(gv.cy+4, gv.cx+12)
		gs.PutCellLoop(gv.cy+5, gv.cx+2)
		gs.PutCellLoop(gv.cy+5, gv.cx+3)
		gs.PutCellLoop(gv.cy+5, gv.cx+4)
		gs.PutCellLoop(gv.cy+5, gv.cx+8)
		gs.PutCellLoop(gv.cy+5, gv.cx+9)
		gs.PutCellLoop(gv.cy+5, gv.cx+10)
		gs.PutCellLoop(gv.cy+7, gv.cx+2)
		gs.PutCellLoop(gv.cy+7, gv.cx+3)
		gs.PutCellLoop(gv.cy+7, gv.cx+4)
		gs.PutCellLoop(gv.cy+7, gv.cx+8)
		gs.PutCellLoop(gv.cy+7, gv.cx+9)
		gs.PutCellLoop(gv.cy+7, gv.cx+10)
		gs.PutCellLoop(gv.cy+8, gv.cx+0)
		gs.PutCellLoop(gv.cy+8, gv.cx+5)
		gs.PutCellLoop(gv.cy+8, gv.cx+7)
		gs.PutCellLoop(gv.cy+8, gv.cx+12)
		gs.PutCellLoop(gv.cy+9, gv.cx+0)
		gs.PutCellLoop(gv.cy+9, gv.cx+5)
		gs.PutCellLoop(gv.cy+9, gv.cx+7)
		gs.PutCellLoop(gv.cy+9, gv.cx+12)
		gs.PutCellLoop(gv.cy+10, gv.cx+0)
		gs.PutCellLoop(gv.cy+10, gv.cx+5)
		gs.PutCellLoop(gv.cy+10, gv.cx+7)
		gs.PutCellLoop(gv.cy+10, gv.cx+12)
		gs.PutCellLoop(gv.cy+12, gv.cx+2)
		gs.PutCellLoop(gv.cy+12, gv.cx+3)
		gs.PutCellLoop(gv.cy+12, gv.cx+4)
		gs.PutCellLoop(gv.cy+12, gv.cx+8)
		gs.PutCellLoop(gv.cy+12, gv.cx+9)
		gs.PutCellLoop(gv.cy+12, gv.cx+10)
	case 10: // Pentadecathlon
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+1, gv.cx+1)
		gs.PutCellLoop(gv.cy+2, gv.cx+0)
		gs.PutCellLoop(gv.cy+2, gv.cx+2)
		gs.PutCellLoop(gv.cy+3, gv.cx+1)
		gs.PutCellLoop(gv.cy+4, gv.cx+1)
		gs.PutCellLoop(gv.cy+5, gv.cx+1)
		gs.PutCellLoop(gv.cy+6, gv.cx+1)
		gs.PutCellLoop(gv.cy+7, gv.cx+0)
		gs.PutCellLoop(gv.cy+7, gv.cx+2)
		gs.PutCellLoop(gv.cy+8, gv.cx+1)
		gs.PutCellLoop(gv.cy+9, gv.cx+1)
	// Spaceships
	case 11: // Glider
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+1, gv.cx+2)
		gs.PutCellLoop(gv.cy+2, gv.cx+0)
		gs.PutCellLoop(gv.cy+2, gv.cx+1)
		gs.PutCellLoop(gv.cy+2, gv.cx+2)
	case 12: // LWSS
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+0, gv.cx+2)
		gs.PutCellLoop(gv.cy+0, gv.cx+3)
		gs.PutCellLoop(gv.cy+0, gv.cx+4)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+4)
		gs.PutCellLoop(gv.cy+2, gv.cx+4)
		gs.PutCellLoop(gv.cy+3, gv.cx+0)
		gs.PutCellLoop(gv.cy+3, gv.cx+3)
	case 13: // MWSS
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+0, gv.cx+2)
		gs.PutCellLoop(gv.cy+0, gv.cx+3)
		gs.PutCellLoop(gv.cy+0, gv.cx+4)
		gs.PutCellLoop(gv.cy+0, gv.cx+5)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+5)
		gs.PutCellLoop(gv.cy+2, gv.cx+5)
		gs.PutCellLoop(gv.cy+3, gv.cx+0)
		gs.PutCellLoop(gv.cy+3, gv.cx+4)
		gs.PutCellLoop(gv.cy+4, gv.cx+2)
	case 14: // HWSS
		gs.PutCellLoop(gv.cy+0, gv.cx+1)
		gs.PutCellLoop(gv.cy+0, gv.cx+2)
		gs.PutCellLoop(gv.cy+0, gv.cx+3)
		gs.PutCellLoop(gv.cy+0, gv.cx+4)
		gs.PutCellLoop(gv.cy+0, gv.cx+5)
		gs.PutCellLoop(gv.cy+0, gv.cx+6)
		gs.PutCellLoop(gv.cy+1, gv.cx+0)
		gs.PutCellLoop(gv.cy+1, gv.cx+6)
		gs.PutCellLoop(gv.cy+2, gv.cx+6)
		gs.PutCellLoop(gv.cy+3, gv.cx+0)
		gs.PutCellLoop(gv.cy+3, gv.cx+5)
		gs.PutCellLoop(gv.cy+4, gv.cx+2)
		gs.PutCellLoop(gv.cy+4, gv.cx+3)
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
			if gs.cell[i][j] >= 1 {
				if buffer[i][j] == 2 || buffer[i][j] == 3 {
					gs.cell[i][j]++
				} else {
					gs.cell[i][j] = 0
				}
			} else {
				if buffer[i][j] == 3 {
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
			if i == 0 && j == 0 {
				continue
			}

			ny := (py + i + GAMEY) % GAMEY
			nx := (px + j + GAMEX) % GAMEX

			if gs.cell[ny][nx] >= 1 {
				n++
			}
		}
	}

	return n
}
