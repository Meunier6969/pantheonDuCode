package main

import (
	rd "math/rand"
)

type GameState struct {
	cell [GAMEY][GAMEY]int
}

func (gf GameState) RandomInitState() {
	for i := 0; i < GAMEY; i++ {
		for j := 0; j < GAMEX; j++ {
			gf.cell[i][j] = rd.Intn(2)
		}
	}
}

func (gf GameState) ExampleInitState() {
	for i := 0; i < GAMEY; i++ {
		for j := 0; j < GAMEX; j++ {
			if (i+j)%2 == 0 {
				gf.cell[i][j] = 1
			}
		}
	}
}