package main

import (

	nc "github.com/rthornton128/goncurses"
)

// func (scr *nc.Window) CustomBorder(gv GameView) { // what couldve been...... 😔
func CustomBorder(scr *nc.Window, gv GameView) {
	var indicator rune

	// Top line
	if gv.vy <= 0 {indicator = '-'} else {indicator = '^'}
	for i := gv.p1x; i <= gv.p2x; i++ {
		scr.MoveAddChar(gv.p1y-1, i, nc.Char(indicator))
	}
	
	// Bottom line
	if gv.vy + (gv.p2y - gv.p1y) < GAMEY - 1 {indicator = 'v'} else {indicator = '-'}
	for i := gv.p1x; i <= gv.p2x; i++ {
		scr.MoveAddChar(gv.p2y+1, i, nc.Char(indicator))
	}
	
	// Left line
	if gv.vx == 0 {indicator = '|'} else {indicator = '<'}
	for i := gv.p1y; i <= gv.p2y; i++ {
		scr.MoveAddChar(i, gv.p1x-1, nc.Char(indicator))
	}
	
	// Right line
	// log.Fatal(gv.vx + (gv.p2x - gv.p1x), GAMEX)
	if gv.vx + (gv.p2x - gv.p1x) < GAMEX - 1 {indicator = '>'} else {indicator = '|'}
	for i := gv.p1y; i <= gv.p2y; i++ {
		scr.MoveAddChar(i, gv.p2x+1, nc.Char(indicator))
	}
}

func PrintInformation(scr *nc.Window, gv GameView, y int, x int) {
	var run string
	if gv.run {
		run = "running"
	} else {
		run = "paused"
	}
	scr.MovePrintf(y,   x, "[c] clear screen")
	scr.MovePrintf(y+1, x, "[r] random state")
	scr.MovePrintf(y+2, x, "[p] %s  ", run)
	scr.MovePrintf(y+3, x, "[s] %d  ", gv.spd)
	scr.MovePrintf(y+4, x, "[hjkl] move around the screen")

	scr.MovePrintf(y+6, x, "[q] quit")
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
}
