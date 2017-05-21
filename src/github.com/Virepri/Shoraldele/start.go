package main

import(
	//"github.github.com/Virepri/Shoraldele/FileIO"
	"github.com/Virepri/Shoraldele/GlobalVars"
	"github.com/jonvaldes/termo"
	"runtime"
	"os"
	"os/exec"
	"fmt"
	"github.com/Virepri/Shoraldele/input"
	"github.com/Virepri/Shoraldele/Display"
	"github.com/Virepri/Shoraldele/FileIO"
	"github.com/Virepri/Shoraldele/Buffer"
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
	//Don't un-defer this! needs to be defered so termo.Stop runs so that user's terminal doesn't get messed up upon exit
	defer os.Exit(0)
	defer exec.Command("reset")
	exec.Command("reset")
	GlobalVars.ConfigLocs = map[string]string{
		"input":"",
		"display":"",
	} //basically add your config location here.
	GlobalVars.SetupFuncs = map[string]func(string) {
		"input":input.Setup,
		"display":display.Dummy,
	} //basically add your setup functions here. the input is meant to be a config location.
	GlobalVars.ModuleRoutines = map[string]func() {
		"input":input.Routine,
		"display":display.DisplayInit,
	} //add your goroutine function here. This should NOT stop until you recieve a "stop" command.

	if err := termo.Init(); err != nil {
		panic(err)
	}
	defer termo.Stop()
	runtime.GOMAXPROCS(len(GlobalVars.ModuleRoutines))

	if len(os.Args) >= 2 {
		data, err := FileIO.Read(os.Args[1])

		if err != nil {
			fmt.Println("Herpderp! Can't read!")
			os.Exit(6)
		}

		buffer.Insert(0, data)
	}

	for k,v := range GlobalVars.SetupFuncs {
		v(GlobalVars.ConfigLocs[k]) //execute all setup function
		go GlobalVars.ModuleRoutines[k]()
		GlobalVars.WaitGroup.Add(1)
	}

	GlobalVars.WaitGroup.Wait()
}
