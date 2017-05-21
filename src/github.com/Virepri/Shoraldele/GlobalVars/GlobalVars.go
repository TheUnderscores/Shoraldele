package GlobalVars

import "sync"

type Command struct {
	Command, Args string
}

var ConfigLocs map[string]string
var SetupFuncs map[string]func(string) //runs on program startup
var ModuleRoutines map[string]func()
var WaitGroup sync.WaitGroup
var MString string = "Command"
