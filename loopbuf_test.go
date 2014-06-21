package loopbuf

import (
	"testing"
	"fmt"
	"bytes"
)

var _ = fmt.Printf // For debugging; delete when done.


func TestLoopBuf(t *testing.T) {
	buf := NewLoopBuf(7)
	if buf.GetBytes() != nil {
		t.Errorf("buf.GetBytes() != nil")
	}

	buf.FillByte(0x01)
	d := buf.GetBytes()
	if d[0] != 0x01 {
		t.Errorf("d[0] == 0x01 failed")
	}

	if buf.DataSize() != 1 {
		t.Errorf("1: size failed")
	}

	buf.FillByte(0x02)
	if buf.DataSize() != 2 {
		t.Errorf("2: size failed")
	}

	buf.FillBytes([]byte{0x03,0x04,0x05,0x06,0x07})
	if bytes.Compare(buf.GetBytes(), []byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07}) != 0 {
		t.Errorf("FillBytes result not right")
	}

	buf.FillByte(0x08)
	expect := buf.GetBytes()
	if bytes.Compare(expect, []byte{0x02,0x03,0x04,0x05,0x06,0x07,0x08}) != 0 {
		t.Errorf("2:%v. %v",expect, buf.Dump())
	}

	buf.FillBytes([]byte{0x99,0x99,0x99,0x99})
	expect = buf.GetBytes()
	if bytes.Compare(expect, []byte{0x06,0x07,0x08,0x99,0x99,0x99,0x99}) != 0 {
		t.Errorf("3:%v. %v",expect, buf.Dump())
	}

	buf.FillBytes([]byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07})
	expect = buf.GetBytes()
	if bytes.Compare(expect, []byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07}) != 0 {
		t.Errorf("3:%v. %v",expect, buf.Dump())
	}

	buf.FillBytes([]byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08})
	expect = buf.GetBytes()
	if bytes.Compare(expect, []byte{0x02,0x03,0x04,0x05,0x06,0x07,0x08}) != 0 {
		t.Errorf("3:%v. %v",expect, buf.Dump())
	}

	// fmt.Println("size:", buf.DataSize())
}

