package main

import (
	// "math"

	nc "github.com/rthornton128/goncurses"
	"log"
)

type GameView struct {
	p1y int // Postion of the top-left corner of the pad on screen
	p1x int
	p2y int // Postion of the bottom-right corner of the pad on screen
	p2x int
	pad *nc.Pad

	vy  int // Position of the view
	vx  int

	cx 	int
	cy 	int

	spd int
	run bool
	cur	bool
	pat int
}

func NewGameView(p1y int, p1x int, p2y int, p2x int) GameView {
	gv := GameView{}

	// Parameters
	gv.p1y = p1y
	gv.p1x = p1x
	gv.p2y = p2y
	gv.p2x = p2x

	// Default values
	gv.vy = 0
	gv.vx = 0
	gv.cy = 0
	gv.cx = 0
	gv.spd = 50
	gv.run = true
	gv.cur = false
	gv.pat = 0

	// Pad thai j'ai faim
	var err error
	gv.pad, err = nc.NewPad(GAMEY, GAMEX)
	if err != nil {
		log.Fatal(err)
	}

	return gv
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

func (gv *GameView) MoveCursor(dir int) {
	switch dir {
	case 8:
		gv.cy -= 1
		gv.cy = Clamp(gv.cy, gv.vy, gv.vy+(gv.p2y-gv.p1y+1))
	case 6:
		gv.cx += 1
		gv.cx = Clamp(gv.cx, gv.vx, gv.vx+(gv.p2x-gv.p1x+1))
	case 4:
		gv.cx -= 1
		gv.cx = Clamp(gv.cx, gv.vx, gv.vx+(gv.p2x-gv.p1x+1))
	case 2:
		gv.cy += 1
		gv.cy = Clamp(gv.cy, gv.vy, gv.vy+(gv.p2y-gv.p1y+1))
	}
}

func (gv *GameView) ToggleCursor() {
	gv.cur = !gv.cur
}

func (gv *GameView) ToggleRun() {
	gv.run = !gv.run
}

func (gv *GameView) ChangeSpeed() {
	gv.spd = 200
}

func (gv *GameView) CyclePattern(dir bool) {
	if dir {
		gv.pat = (gv.pat+1 + 15) % 15
	} else {
		gv.pat = (gv.pat-1 + 15) % 15
	}
}

func (gv *GameView) DrawCursor()  {
	gv.pad.MoveAddChar(gv.cy, gv.cx, gv.pad.MoveInChar(gv.cy, gv.cx) | nc.A_STANDOUT)
}

func (gv GameView) NoutRefresh(scr *nc.Window) {
	gv.pad.NoutRefresh(gv.vy, gv.vx, gv.p1y, gv.p1x, gv.p2y, gv.p2x)
	// gv.pad.NoutRefresh(0, 0, 5, 5, 10, 10)
}
