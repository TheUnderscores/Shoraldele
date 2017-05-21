package buffer

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
	overwrite(offset, data)
}

func Delete(offset int, length int) {
	work_buffer = append(work_buffer[:offset], work_buffer[offset+len(data):])
}

func GetBufferContents(offset int, length int) {
	if length < 1 {
		return work_buffer[offset:]
	}

	return work_buffer[offset:offset + length]
}
