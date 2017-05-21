package input

//input manages modes

import (
	tm "github.com/jonvaldes/termo"
	gv "github.com/Virepri/Shoraldele/GlobalVars"
	"github.com/Virepri/Shoraldele/Buffer"
	"github.com/Virepri/Shoraldele/Codes"
	"reflect"
)

//http://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go

type ModeType uint8

const (
	Command ModeType = iota
	Insert
	Select
)

var CurrentMode ModeType //Defaults to 0, which is Command mode

//Append to this array if you want to be notified when the mode changes
var ModeChangeNotifiers [](chan<- ModeType)

var HasSelection bool

func Setup(_ string) {
	return
}

func Routine() {
	defer gv.WaitGroup.Done()

	keychan := make(chan tm.ScanCode)
	errchan := make(chan error)

	tm.StartKeyReadLoop(keychan, errchan)

	for {
		handleKey(<-keychan)
	}
	return
}

func changeMode (new_mode ModeType) {
	if new_mode == CurrentMode {
		// do nothing!
		return
	} else {
		CurrentMode = new_mode
		// TODO: This should probably dispatch threads so that someone adding a chan but not listening does not block
		for _, chn := range ModeChangeNotifiers {
			chn <- new_mode
		}
	}
}

func handleKey (keycode tm.ScanCode) {
	if s(keycode, codes.ESC) { //ESC
		changeMode(Command)
		return
	}
	switch {
	case s(keycode, codes.Ci): //i
		changeMode(Insert)
	case s(keycode, codes.Cs): //s
		changeMode(Select)
	case s(keycode, codes.Cc): //c
		changeMode(Command)

	case s(keycode, codes.LEFT):
		if buffer.GetCursorPosition() > 0 {
			buffer.SetCursorPosition(buffer.GetCursorPosition() + 1)
		}
	case s(keycode, codes.RIGHT):
		if buffer.GetCursorPosition() < buffer.GetBufferSize() {
			buffer.SetCursorPosition(buffer.GetCursorPosition() - 1)
		}
	case s(keycode, codes.UP):
		line, char, lines := buffer.GetCursorLinePosition()
		if line > 0 {
			prevline := lines[line - 1]
			buffer.SetCursorPosition(prevline[0] + min(prevline[1], char))
		}
	}
}	

func min (a, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}
		

func s (a tm.ScanCode, b tm.ScanCode) bool {
	return reflect.DeepEqual(a, b)
}
