package buffer

import(
	"github.com/Virepri/Shoraldele/GlobalVars"
	"fmt"
)

func Init(ini string) {
	fmt.Println("buffer.Init: STUB")
}

func Entry() {
	fmt.Println("buffer.Entry: STUB")
	GlobalVars.WaitGroup.Done()
}

func PushCommand(cmd, args string) {
	fmt.Println("buffer.PushCommand: STUB")
}
