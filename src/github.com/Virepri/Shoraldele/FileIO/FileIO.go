package FileIO

import (
	"os"
	"github.com/Virepri/Shoraldele/Buffer"
)

var file os.File

func Read(name string) (string, error) {
	if fi, err := os.OpenFile(name, os.O_RDWR | os.O_CREATE, 0755); err == nil {
		finfo,_ := fi.Stat()
		dat := make([]byte, finfo.Size())
		fi.Read(dat)
		return string(dat), nil
	} else {
		return "", err
	}
}

func Write() {
	dat := buffer.GetBufferContents(0,-1)
	file.Truncate(0)
	file.Sync()
	file.Write(dat)
}