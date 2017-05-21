package display

import (
	"github.com/Virepri/Shoraldele/Buffer"
	"github.com/jonvaldes/termo"
	"fmt"
)

var running bool = true

func StopDisplay() {
	running = false
}

func DisplayInit() {
	tsx,tsy,_ := termo.Size()
	buffer.Insert(0,fmt.Sprint(tsx,tsy))
	framebuffer := termo.NewFramebuffer(tsx, tsy)

	var State termo.CellState

	State.Attrib = 0
	State.FGColor = 30
	State.BGColor = 80

	//termo.Init() //Init stuff
	for running {
		//framebuffer.ASCIIRect(0, 0, 80, 20, true, false)//Here we draw an ascii rectangle of size 40x40
		framebuffer.ASCIIRect(0,0,tsx-1,tsy-1,true,false)

		framebuffer.AttribText(1, 1, State, string(buffer.GetBufferContents(0, -1)))

		framebuffer.Flush()

		framebuffer.Clear()
	}
}

func Dummy(_ string){

}
