package main

import(
	"github.com/Virepri/Shoraldele/GlobalVars"
	"github.com/jonvaldes/termo"
	"runtime"
	"os"
	"github.com/Virepri/Shoraldele/input"
)

/*
ADDING YOUR MODULE:
You'll need 3 functions at a minimum, as well as a config location (can be empty)
Setup: Recieves a string. Reads the config, sets up.
Commands: Capitalize anything you want others to use. Document them at the top of your file.
Routine: Recieves nothing. Should _ALWAYS_ be running until you recieve the "stop" command.

Basically, add a new line to each of the below maps with this setup:
"<module name>":<what it wants>,

and then, to the import statement, add the directory path to your module, ignoring src/
*/

func main(){
	GlobalVars.ConfigLocs = map[string]string{
		"input":"",
	} //basically add your config location here.
	GlobalVars.SetupFuncs = map[string]func(string) {
		"input":input.Setup,
	} //basically add your setup functions here. the input is meant to be a config location.
	GlobalVars.ModuleRoutines = map[string]func() {
		"input":input.Routine,
	} //add your goroutine function here. This should NOT stop until you recieve a "stop" command.

	if err := termo.Init(); err != nil {
		panic(err)
	}
	defer termo.Stop()
	runtime.GOMAXPROCS(len(GlobalVars.ModuleRoutines))

	for k,v := range GlobalVars.SetupFuncs {
		v(GlobalVars.ConfigLocs[k]) //execute all setup functions
		go GlobalVars.ModuleRoutines[k]()
		GlobalVars.WaitGroup.Add(1)
	}

	GlobalVars.WaitGroup.Wait()

	os.Exit(0)
}
