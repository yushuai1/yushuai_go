package yuaa

import (
	"encoding/binary"
	"unsafe"
)

type My struct {
	Age int
	Name string
}


type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func StructToBytes(s *My) []byte {
	Len := unsafe.Sizeof(*s)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(s)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))
	return data
}

func BytesToStruct(data []byte) *My  {
	return *(**My)(unsafe.Pointer(&data))
}

func UInt32ToBytes(i int) []byte {
	var buf = make([]byte,4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToUInt32(buf []byte) int {
	return int(binary.BigEndian.Uint32(buf))
}
