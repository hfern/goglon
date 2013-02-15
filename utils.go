package goglon

import (
	"bufio"
	"fmt"
)

func checkNextByte(reader *bufio.Reader, b byte, g *byte) bool {
	nextByte, _ := reader.Peek(1)
	*g = nextByte[0]
	if *g != b {
		return false
	}
	return true
}

func assertNextByte(reader *bufio.Reader, b byte) {
	var g byte
	if !checkNextByte(reader, b, &g) {
		fmtstr := "Bad byte assertion. Expected %#x, got %#x."
		panic(fmt.Sprintf(fmtstr, b, g))
	}
	reader.ReadByte()
	return
}
