package main
/*
TODO
single step

cursor
add shapes

*/
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
	// Checking for terminal size

	vertical := 2*wy >= wx
	var offsety int
	var offsetx int
	if vertical {
		offsety = 9
		offsetx = 3
	} else {
		offsety = 3
		offsetx = 36
	}

	var gs GameState
	gs.RandomInitState()

	gv := GameView{
		0,
		0,
		2,
		3,
		Min(102, wy - offsety),
		Min(103, wx - offsetx),
		nil, // placeholder for 2 lines below
		200,
		true,
	}
	gv.pad, err = nc.NewPad(GAMEY, GAMEX)
	if err != nil {
		log.Fatal(err)
	}

	for {	
		// Actual drawing
		CustomBorder(scr, gv)
		if gv.run {
			gs.ComputeNextStep()
		}
		gv.FillWithState(gs)

		if vertical {
			PrintInformation(scr, gv, wy-6, 3)
		} else {
			PrintInformation(scr, gv, 2, gv.p2x + 4)
		}

		// Refresh screen
		scr.NoutRefresh()   // Border and ui
		gv.NoutRefresh(scr) // Game of Life

		nc.Update()

		// User Input
		ch := scr.GetChar()
		switch ch {
		// Yeah
		case 'c':
			gs.EmptyInitState()
		case 'r':
			gs.RandomInitState()

		case 'q':
			return
		case 'p':
			gv.ToggleRun()
		// Moving
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
