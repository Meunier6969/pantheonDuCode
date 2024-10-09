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

	scr.Timeout(50)

	wy, wx := scr.MaxYX()
	if wy < 18 || wx < 52 { // "fuck it we ball numbers"
		log.Fatal("Terminal is not big enough")
	}

	gv := GameView{
		0,
		0,
		5,
		5,
		25,
		50,
		nil, // placeholder for 2 lines below
	}
	gv.pad, err = nc.NewPad(GAMEY, GAMEX)
	if err != nil {
		log.Fatal(err)
	}

	// Fit to screen
	

	var ch nc.Key
	for {
		// Actual drawing
		scr.MovePrintf(0, 0, "-- %d, %d -- %v --", gv.vy, gv.vx, ch)
		scr.MovePrintf(1, 0, "-- %d, %d -- %d, %d -- ", gv.p1y, gv.p1y, gv.p2y, gv.p2y)
		gv.TestFilling()
		CustomBorder(scr, gv)
		// Draw to screen
		scr.NoutRefresh() // Border and ui
		gv.NoutRefresh(scr) // Game of Life

		nc.Update()

		ch := gv.pad.GetChar()
		switch ch {
			case 'q':
				return
			case 'h':
				gv.MoveView(4)
			case 'j':
				gv.MoveView(2)
			case 'k':
				gv.MoveView(8)
			case 'l':
				gv.MoveView(6)
		}
	}
}
