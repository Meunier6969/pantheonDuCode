package main

import (
	// "math"

	nc "github.com/rthornton128/goncurses"
	// "log"
)

type GameView struct {
	vy  int // Position of the view
	vx  int
	p1y int // Postion of the top-left corner of the pad on screen
	p1x int
	p2y int // Postion of the bottom-right corner of the pad on screen
	p2x int
	pad *nc.Pad

	spd int
	run bool
}

func (gv GameView) TestFilling() {
	for i := 0; i < GAMEY; i++ {
		for j := 0; j < GAMEX; j++ {
			if (i * j % 3) == 1 {
				gv.pad.MoveAddChar(i, j, '@')
			}
		}
	}
}

func (gv *GameView) FillWithState(gs GameState) {
	for i := 0; i < GAMEY; i++ {
		for j := 0; j < GAMEX; j++ {
			if gs.cell[i][j] > 0 {
				gv.pad.MoveAddChar(i, j, '#')
			} else {
				gv.pad.MoveAddChar(i, j, ' ')
			}
		}
	}
}

func (gv *GameView) MoveView(dir int) {
	switch dir {
	case 8:
		gv.vy -= 1
		gv.vy = Max(0, gv.vy)
	case 6:
		gv.vx += 1
		gv.vx = Min(100-(gv.p2x-gv.p1x+1), gv.vx)
	case 4:
		gv.vx -= 1
		gv.vx = Max(0, gv.vx)
	case 2:
		gv.vy += 1
		gv.vy = Min(100-(gv.p2y-gv.p1y+1), gv.vy)
	}
}

func (gv *GameView) ToggleRun() {
	gv.run = !gv.run
}

func (gv *GameView) ChangeSpeed() {
	gv.spd = 200
}

func (gv GameView) NoutRefresh(scr *nc.Window) {
	gv.pad.NoutRefresh(gv.vy, gv.vx, gv.p1y, gv.p1x, gv.p2y, gv.p2x)
	// gv.pad.NoutRefresh(0, 0, 5, 5, 10, 10)
}
