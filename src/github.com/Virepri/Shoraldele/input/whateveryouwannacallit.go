package input

//input manages modes

import (
	tm "github.com/jonvaldes/termo"
	gv "github.com/Virepri/Shoraldele/GlobalVars"
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

func Setup(_ string) {
}

func Command(cmd string) {
	switch(cmd){

	}
}

func Routine {
	defer gv.WaitGroup.Done()

	
	keychan := make(chan tm.ScanCode)
	errchan := make(chan error)

	tm.StartKeyReadLoop(keychan, errchan)

	var keycode tm.ScanCode
	for {
		keycode = <-keychan
		
	}
}

func changeMode (new_mode ModeType) {
	if new_mode == CurrentMode {
		// do nothing!
		return
	} else {
		CurrentMode = new_mode
		// TODO: This should probably dispatch threads so that someone adding a chan but not listening does not block
		for _, chn = range ModeChangeNotifiers {
			chn <- new_mode
		}
	}
}
