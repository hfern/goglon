package goglon

import (
	"bufio"
)

func _TableEncode(writer *bufio.Writer, dtype interface{}) {
	return
}

func _StringEncode(writer *bufio.Writer, dtype interface{}) {
	return
}

func _NumberEncode(writer *bufio.Writer, dtype interface{}) {
	return
}

func _ArrayEncode(writer *bufio.Writer, dtype interface{}) {
	return
}

func _FalseEncode(writer *bufio.Writer, dtype interface{}) {
	writer.WriteByte(BYTE_FALSE)
	return
}

func _TrueEncode(writer *bufio.Writer, dtype interface{}) {
	writer.WriteByte(BYTE_TRUE)
	return
}
