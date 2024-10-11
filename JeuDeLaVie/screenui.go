package main

import (
	nc "github.com/rthornton128/goncurses"
)

// func (scr *nc.Window) CustomBorder(gv GameView) { // what couldve been...... 😔
func CustomBorder(scr *nc.Window, gv GameView) {
	var indicator rune

	// Top line
	if gv.vy <= 0 {
		indicator = '-'
	} else {
		indicator = '^'
	}
	for i := gv.p1x; i <= gv.p2x; i++ {
		scr.MoveAddChar(gv.p1y-1, i, nc.Char(indicator))
	}

	// Bottom line
	if gv.vy+(gv.p2y-gv.p1y) < GAMEY-1 {
		indicator = 'v'
	} else {
		indicator = '-'
	}
	for i := gv.p1x; i <= gv.p2x; i++ {
		scr.MoveAddChar(gv.p2y+1, i, nc.Char(indicator))
	}

	// Left line
	if gv.vx == 0 {
		indicator = '|'
	} else {
		indicator = '<'
	}
	for i := gv.p1y; i <= gv.p2y; i++ {
		scr.MoveAddChar(i, gv.p1x-1, nc.Char(indicator))
	}

	// Right line
	// log.Fatal(gv.vx + (gv.p2x - gv.p1x), GAMEX)
	if gv.vx+(gv.p2x-gv.p1x) < GAMEX-1 {
		indicator = '>'
	} else {
		indicator = '|'
	}
	for i := gv.p1y; i <= gv.p2y; i++ {
		scr.MoveAddChar(i, gv.p2x+1, nc.Char(indicator))
	}
}

func DrawInformationSide(scr *nc.Window, gv GameView, y int, x int) {

	var cursortoggled string

	if gv.cur {
		cursortoggled = "on"
	} else {
		cursortoggled = "off"
	}

	// Clear
	for i := 0; i < 12; i++ {
		scr.MovePrintf(y+i, x, "                                       ")
	}

	scr.MovePrintf(y, x, "[q] quit")

	scr.MovePrintf(y+2, x, "[r] random state")
	scr.MovePrintf(y+3, x, "[c] clear screen")

	scr.MovePrintf(y+5, x, "[p] pause")
	if gv.run {
		scr.MovePrintf(y+6, x, "-- RUNNING --")

		scr.MovePrintf(y+8, x, "[s] change speed")
		scr.MovePrintf(y+9, x, "[hjkl] move around the screen")

	} else {
		scr.MovePrintf(y+6, x, "-- PAUSED -- ")

		scr.MovePrintf(y+8, x, "[s] single step")
		scr.MovePrintf(y+9, x, "[t] toggle cursor (%s) ", cursortoggled)
		if gv.cur {
			scr.MovePrintf(y+10, x, "[hjkl] move around the cursor")
			scr.MovePrintf(y+11, x, "[SPACE] toggle selected cell")
			
		} else {
			scr.MovePrintf(y+10, x, "[hjkl] move around the screen")
			scr.MovePrintf(y+11, x, "                            ")
		}
	}

}

func DrawInformationBottom(scr *nc.Window, gv GameView, y int, x int) {
	var cursortoggled string

	if gv.cur {
		cursortoggled = "on"
	} else {
		cursortoggled = "off"
	}

	// Clear
	for i := 0; i < 7; i++ {
		scr.MovePrintf(y+i, x+20, "                                      ")
	}

	scr.MovePrintf(y, x, "[q] quit")

	scr.MovePrintf(y+2, x, "[r] random state")
	scr.MovePrintf(y+3, x, "[c] clear screen")

	scr.MovePrintf(y+4, x, "[p] pause")
	if gv.run {
		scr.MovePrintf(y, x+20, "-- RUNNING --")

		scr.MovePrintf(y+2, x+20, "[s] change speed")
		scr.MovePrintf(y+3, x+20, "[hjkl] move around the screen")

	} else {
		scr.MovePrintf(y+0, x+20, "-- PAUSED -- ")

		scr.MovePrintf(y+2, x+20, "[s] single step")
		scr.MovePrintf(y+3, x+20, "[t] toggle cursor (%s) ", cursortoggled)
		if gv.cur {
			scr.MovePrintf(y+4, x+20, "[hjkl] move around the cursor")
			scr.MovePrintf(y+5, x+20, "[SPACE] toggle selected cell")
			
		} else {
			scr.MovePrintf(y+4, x+20, "[hjkl] move around the screen")
			scr.MovePrintf(y+5, x+20, "                            ")
		}
	}

}

/*
[c] clear screen
[r] random state

[p] running
[s] speed
[hjkl] move around the screen

[p] paused
[s] single step
[hjkl] move around the screen/cursor
[?] toggle cursor
*/
