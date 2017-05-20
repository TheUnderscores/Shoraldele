package main

import(
	"github.com/Virepri/Shoraldele/GlobalVars"
	"runtime"
	"os"
)

/*
ADDING YOUR MODULE:
You'll need 3 functions at a minimum, as well as a config location (can be empty)
Setup: Recieves a string. Reads the config, sets up.
Command: Recieves 2 strings. Executes the corresponding command (string 1) with the arguments (string 2)
Routine: Recieves nothing. Should _ALWAYS_ be running until you recieve the "stop" command.

Basically, add a new line to each of the below maps with this setup:
"<module name>":<what it wants>,

and then, to the import statement, add the directory path to your module, ignoring src/
*/

func main(){
	GlobalVars.ConfigLocs = map[string]string{} //basically add your config location here.
	GlobalVars.SetupFuncs = map[string]func(string) {} //basically add your setup functions here. the input is meant to be a config location.
	GlobalVars.CmdFuncs = map[string]func(string,string) {} //add your command function here
	GlobalVars.ModuleRoutines = map[string]func() {} //add your goroutine function here. This should NOT stop until you recieve a "stop" command.

	runtime.GOMAXPROCS(len(GlobalVars.ModuleRoutines))

	for k,v := range GlobalVars.SetupFuncs {
		v(GlobalVars.ConfigLocs[k]) //execute all setup functions
		go GlobalVars.ModuleRoutines[k]()
		GlobalVars.WaitGroup.Add(1)
	}

	GlobalVars.WaitGroup.Wait()
	os.Exit(0)
}
