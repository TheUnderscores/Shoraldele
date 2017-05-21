package buffer

var work_buffer []byte

var cursorPosition int

func Overwrite(offset int, data string) {
	if len(data) + len(work_buffer) > cap(work_buffer) {
		work_buffer = append(work_buffer, make([]byte, 1024)...)
	}

	copy(work_buffer[offset:], []byte(data))
}

func Insert(offset int, data string) {
	if len(data) + len(work_buffer) > cap(work_buffer) {
		work_buffer = append(work_buffer, make([]byte, 1024)...)
	}

	copy(work_buffer[offset+len(data):], work_buffer[offset:])
	Overwrite(offset, data)
}

func Delete(offset int, length int) {
	work_buffer = append(work_buffer[:offset], work_buffer[offset+length:]...)
}

func GetBufferContents(offset int, length int) []byte {
	if length < 1 {
		return work_buffer[offset:]
	}

	return work_buffer[offset:offset + length]
}

func GetBufferSize() int {
	return len(work_buffer)
}

func GetCursorPosition() uint64 {
	return cursorPosition
}

func SetCursorPosition(newpos uint64) {
	cursorPosition = newpos
	//TODO: Call some function to notify output module
}

//Wraps at 80 chars, hardcoded 'cause codeday
//Each element is [length, offset]
func BufferLineSplit() (res [][2]int) {
	res = make([][2]int)
	//The starting position of the line currently being processed
	startpos := 0
	size := GetBufferSize()
	for curpos := 0 ; curpos < size ; curpos++ {
		length = curpos - startpos + 1
		if work_buffer[curpos] == "\n" || length == 80 {
			var a [2]int
			a[0] = curpos - startpos + 1
			a[1] = startpos
			append(res, a)
			startpos = curpos + 1
		}
	}
	return
}


func GetCursorLinePosition() (int, int) {
	l = BufferLineSplit()
	for i, a := range l {
		if a[0] + a[1] >= cursorPosition {
			return i, cursorPosition - a[1]
		}
	}
	return len(work_buffer)
}
