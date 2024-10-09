package main

import (

	nc "github.com/rthornton128/goncurses"
)

// func (scr *nc.Window) CustomBorder(gv GameView) {
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