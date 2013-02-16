package goglon

var typeswitch map[byte]typeMarshall
var Opcodes []byte

func init() {
	typeswitch = map[byte]typeMarshall{
		BYTE_TABLE:  typeMarshall{_TableEncode, _TableDecode},
		BYTE_ARRAY:  typeMarshall{_ArrayEncode, _ArrayDecode},
		BYTE_FALSE:  typeMarshall{_FalseEncode, _FalseDecode},
		BYTE_TRUE:   typeMarshall{_TrueEncode, _TrueDecode},
		BYTE_NUMBER: typeMarshall{_NumberEncode, _NumberDecode},
		BYTE_STRING: typeMarshall{_StringEncode, _StringDecode},
		BYTE_VECTOR: typeMarshall{nil, _VectorDecode},
		BYTE_ANGLE:  typeMarshall{nil, _AngleDecode},
	}

	Opcodes = make([]byte, len(typeswitch))
	i := 0
	for code, _ := range typeswitch {
		Opcodes[i] = code
		i++
	}
}
