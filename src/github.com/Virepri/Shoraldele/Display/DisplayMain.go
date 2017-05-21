package display

import (
	gv "github.com/Virepri/Shoraldele/GlobalVars"
	"github.com/Virepri/Shoraldele/Buffer"
	"github.com/jonvaldes/termo"
	"strings"
)

var running bool = true

func StopDisplay() {
	running = false
}

func DisplayInit() {
	w,h,_ := termo.Size()
	f := termo.NewFramebuffer(w,h)

	var State, BarState termo.CellState

	State.Attrib = 0
	State.FGColor = termo.ColorGreen
	State.BGColor = termo.ColorDefault

	BarState.Attrib = 0
	BarState.FGColor = termo.ColorDefault
	BarState.BGColor = termo.ColorDefault

	f.Clear()
	f.Flush()
	//termo.Init() //Init stuff
	for running {
		if _w, _h, _ := termo.Size(); w != _w || h != _h {
			f.Clear()
			f.ASCIIRect(0,0,w-1,h-1,true,false)
			f.Flush()
			w = _w
			h = _h
			f = termo.NewFramebuffer(w, h)
		}
		f.Clear()

		for i := 0 ; i < w ; i++ {
			f.AttribText(i, 0, BarState, "\u2550")
			f.AttribText(i, h - 2, BarState, "\u2550")
		}

		f.AttribText(1,h-2, State, gv.MString)

		for k,v := range GetWraps(string(buffer.GetBufferContents(0,-1)), w - 2, h - 2) {
			if k <= h - 2 {
				f.AttribText(1,1+k, State, v)
			}
		}

		f.Flush()
	}
}

func Dummy(_ string){

}

func GetWraps(dat string, w, h int) []string {
	o := []string{}
	for _,v := range strings.Split(dat,"\n") {
		o = append(o,SplitNLen(v,w)...)
	}
	return o
}

func SplitNLen(s string, n int) []string {
	tmpstr := s
	o := []string{}
	for int(len(tmpstr) / n) != 0 {
		o = append(o,tmpstr[:n])
		tmpstr = tmpstr[n:]
	}
	if len(tmpstr) != 0 {
		//o = PadRight(o, n,' ')
		o = append(o, PadRight(tmpstr,n,' '))
	}
	return o
}

func PadRight(s string, n int, padding rune) string {
	o := s
	for len(o) != n {
		o += string(padding)
	}
	return o;
}
