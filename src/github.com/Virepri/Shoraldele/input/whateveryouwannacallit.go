package input

import (
	tm "github.com/jonvaldes/termo"
	gv "github.com/Virepri/Shoraldele/GlobalVars"
)
	
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
