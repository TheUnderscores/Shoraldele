package display

import (
	"github.com/Virepri/Shoraldele/Buffer"
	"github.com/jonvaldes/termo"
)



func DisplayInit(z string) {
	framebuffer := termo.NewFramebuffer(80, 20) 
	var State termo.CellState

	State.Attrib = 5
	State.FGColor = 30
	State.BGColor = 80

	termo.Init() //Init stuff
	for {
		
		framebuffer.ASCIIRect(0, 0, 80, 20, true, false)//Here we draw an ascii rectangle of size 40x40

		framebuffer.AttribText(1, 1, State, string(buffer.GetBufferContents(0, -1)))

		framebuffer.Flush()

		framebuffer.Clear()
	}
}

func Dummy(){

}