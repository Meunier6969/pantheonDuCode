package main

import (
	"log"
	// "math"

	nc "github.com/rthornton128/goncurses"
)

const (
	GAMEX int = 100
	GAMEY int = 100
)

func main() {
	scr, err := nc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer nc.End()

	nc.Cursor(0)
	nc.Echo(false)
	nc.Raw(true)

	// Checking terminal size
	wy, wx := scr.MaxYX()

	vertical := 2*wy >= wx
	var offsety int
	var offsetx int
	if vertical {
		offsety = 10
		offsetx = 3
	} else {
		offsety = 3
		offsetx = 36
	}

	var gs GameState
	gs.RandomInitState()

	gv := NewGameView(2, 3, Min(102, wy-offsety), Min(103, wx-offsetx))

	// Game speed
	scr.Timeout(gv.spd)

	// Main loop
	for {
		// Actual drawing
		CustomBorder(scr, gv)
		if gv.run {
			gs.ComputeNextStep()
		}
		gv.FillWithState(gs)

		if gv.cur && !gv.run {
			gv.DrawCursor()
		}

		if vertical {
			DrawInformationBottom(scr, gv, wy-7, 3)
		} else {
			DrawInformationSide(scr, gv, 2, gv.p2x+4)
		}

		// Refresh screen
		scr.NoutRefresh()   // Border and ui
		gv.NoutRefresh(scr) // Game of Life

		nc.Update()

		// User Input
		ch := scr.GetChar()
		switch ch {
		// Yeah
		case 'q':
			return

		case 'c':
			gs.EmptyInitState()
		case 'r':
			gs.RandomInitState()

		case 's':
			if gv.run {
				gv.ChangeSpeed()
			} else {
				gs.ComputeNextStep()
			}
		case 'p':
			gv.ToggleRun()

		// When paused
		case 't':
			gv.ToggleCursor()

		case ' ':
			gs.AddPatternAtCursor(gv)
		case 'a':
			gv.CyclePattern(false)
		case 'z':
			gv.CyclePattern(true)

		// Moving
		case 'h':
			if gv.run || !gv.cur {
				gv.MoveView(4)
			} else {
				gv.MoveCursor(4)
			}
		case 'j':
			if gv.run || !gv.cur {
				gv.MoveView(2)
			} else {
				gv.MoveCursor(2)
			}
		case 'k':
			if gv.run || !gv.cur {
				gv.MoveView(8)
			} else {
				gv.MoveCursor(8)
			}
		case 'l':
			if gv.run || !gv.cur {
				gv.MoveView(6)
			} else {
				gv.MoveCursor(6)
			}
		default:
		}
	}
}
