package loopbuf
import (
	"fmt"
	//"io"
)
var _ = fmt.Printf

type LoopBuf struct {
	buf []byte
	bufSize int
	head int	//point to the next available data
	next int
	dataSize int
}


func NewLoopBuf(size int) *LoopBuf{
	s := new(LoopBuf)
	s.buf = make([]byte, size)
	s.bufSize = size
	s.head = 0
	s.next = 0
	s.dataSize = 0
	return s
}


func (s *LoopBuf) FillByte(c byte) {
	if s.dataSize >= s.bufSize {

	}
	// fmt.Printf("FillByte:%x  idx:%d\n", c, s.next)
	s.buf[s.next] = c
	s.next = (s.next+1)%s.bufSize
	if s.dataSize < s.bufSize {
		s.dataSize ++
	} else {
		s.head = (s.head + 1)%s.bufSize
	}
}

func (s *LoopBuf) FillBytes(bs []byte) {
	for _,v := range bs {
		s.buf[s.next] = v
		s.next = (s.next+1)%s.bufSize
		s.dataSize ++
	}
	s.next = s.next%s.bufSize
	if s.dataSize > s.bufSize {
		ovflow := s.dataSize - s.bufSize
		s.dataSize = s.bufSize
		s.head = (s.head + ovflow)%s.bufSize
	}
	// fmt.Println(s.next)
}

func (s *LoopBuf) GetBytes() []byte {
	if (s.dataSize == 0) {
		return nil
	}
	return append(s.buf[s.head:], s.buf[0:s.head]...)
}

func (s *LoopBuf) DataSize() int {
	return s.dataSize
}

func (s *LoopBuf) Dump() string {
	return fmt.Sprintf("head:%v | next:%v | dataSize:%v", s.head, s.next, s.dataSize)
}

