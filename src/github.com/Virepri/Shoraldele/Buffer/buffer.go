package buffer

import "strings"

var work_buffer []byte

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

func Replace(torep, repwith string) {
	t := string(work_buffer)

	strings.Replace(t,torep,repwith,-1)

	work_buffer = []byte(t)
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
