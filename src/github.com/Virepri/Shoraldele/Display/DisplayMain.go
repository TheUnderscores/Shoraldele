package display

import (
	"github.com/Virepri/Shoraldele/Buffer"
	"github.com/jonvaldes/termo"
)



func DisplayInit(z string) {
	framebuffer := termo.NewFramebuffer(80, 20) 
	var fill string
	var State termo.CellState

	fill = "abcdefgijklmnopqrstuvwxyz"

	State.Attrib = 0
	State.FGColor = 30
	State.BGColor = 80

	buffer.Insert(1, fill)

	termo.Init() //Init stuff
	for i := 0; i < 10; i++ {
		
		framebuffer.ASCIIRect(0, 0, 80, 20, true, false)//Here we draw an ascii rectangle of size 40x40

		framebuffer.AttribText(1, 1, State, string(buffer.GetBufferContents(0, 20)))
		buffer.Insert(1, string(i))
		framebuffer.Flush()

		framebuffer.Clear()
	}
}

func Dummy(){

}