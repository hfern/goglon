package goglon

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Decode(reader *bufio.Reader) (interface{}, error) {
	//buf := bufio.NewReader(*reader)
	return _unserialize(reader)
}

func DecodeFromFile(file *os.File) (interface{}, error) {
	return Decode(bufio.NewReader(file))
}

func _unserialize(reader *bufio.Reader) (interface{}, error) {
	// read type byte
	bt, _ := reader.Peek(1)
	typ := bt[0]

	marshallers, ok := typeswitch[typ]

	if !ok {
		fmtstr := "Unknown opcode (%#x). Input may be corrupt."
		return nil, errors.New(fmt.Sprintf(fmtstr, typ))
	}

	val := marshallers.decoder(reader)

	return val, nil
}

func _TableDecode(in *bufio.Reader) interface{} {
	assertNextByte(in, 0x2)
	tbl := TTable{}

	for {
		bt, err := in.Peek(1)
		opcode := bt[0]
		if err != nil {
			return tbl
		}
		// last opcode to end is reached
		if opcode == BYTE_END {
			in.ReadByte() // advance pointer to consume peek opcode
			return tbl
		}

		key, err := _unserialize(in)
		if err != nil {
			panic(err)
		}

		val, err := _unserialize(in)
		if err != nil {
			panic(err)
		}

		tbl[key] = val
	}

	return tbl
}

func _StringDecode(in *bufio.Reader) interface{} {
	assertNextByte(in, 0x7)

	str := bytes.Buffer{}
	escape := false
	for {
		character, err := in.ReadByte()
		if err == io.EOF {
			panic("Expected unescaped \\1 to end string! (Got EOF)")
		} else if escape {
			// handle escape condition for next character
			if character == 0x3 {
				str.WriteByte(0x0)
			} else {
				str.WriteByte(character)
			}
		} else if character == 0x2 {
			escape = true
		} else if character == BYTE_END {
			return str.String()

		} else {
			str.WriteByte(character)
		}
	}
	return str.String()
}

func _NumberDecode(in *bufio.Reader) interface{} {
	assertNextByte(in, 0x6)

	num := bytes.Buffer{} // stored as string within the text file

	for {
		character, err := in.ReadByte()

		if err == io.EOF {
			panic("Expected unescaped \\1 to end number! (Got EOF)")
		} else if character == BYTE_END {
			break
		} else {
			num.WriteByte(character)
		}
	}

	byteified := num.Bytes()
	isFloat := false

	for _, byteval := range byteified {
		if byteval == '.' {
			isFloat = true
			break
		}
	}

	if isFloat {
		fl, err := strconv.ParseFloat(string(byteified), 64)
		if err == nil {
			return fl
		} else {
			return float64(0)
		}
	}

	intval, err := strconv.Atoi(string(byteified))
	if err == nil {
		return intval
	}

	return int64(0)
}

func _ArrayDecode(in *bufio.Reader) interface{} {
	assertNextByte(in, 0x3)
	arr := make(TArray, 0, 0)

	for {
		peeks, err := in.Peek(1)
		if err == io.EOF {
			panic("Early EOF!")
		}

		opcode := peeks[0]
		if opcode == BYTE_END {
			in.ReadByte()
			return arr
		} else {
			val, err := _unserialize(in)
			if err != nil {
				panic(err)
			}
			arr = append(arr, val)
		}
	}
	return arr
}
