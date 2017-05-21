package FileIO

import (
	"os"
	"github.com/Virepri/Shoraldele/Buffer"
)

var file *os.File

func Read(name string) (string, error) {
	var err error
	if file, err = os.OpenFile(name, os.O_RDWR | os.O_CREATE, 0644); err == nil {
		filenfo,_ := file.Stat()
		dat := make([]byte, filenfo.Size())
		file.Read(dat)
		return string(dat), nil
	} else {
		return "", err
	}
}

func Write() {
	dat := buffer.GetBufferContents(0,-1)
	file.Truncate(0)
	file.WriteAt(dat,0)
}
