package goglon

import (
	"bufio"
)

type typeMarshall struct {
	encoder func(*bufio.Writer, interface{})
	decoder func(*bufio.Reader) interface{}
}

const (
	BYTE_END         byte = 0x1
	BYTE_TABLE            = 0x2
	BYTE_ARRAY            = 0x3
	BYTE_FALSE            = 0x4
	BYTE_TRUE             = 0x5
	BYTE_NUMBER           = 0x6
	BYTE_STRING           = 0x7
	BYTE_VECTOR           = 0x8
	BYTE_ANGLE            = 0x9
	BYTE_ENTITY           = 0xA
	BYTE_PLAYER           = 0xB
	BYTE_CEFFECTDATA      = 0xC
	BYTE_CONVAR           = 0xD
	BYTE_COLOR            = 0xF
)

/**
 * See http://stronghold.googlecode.com/svn/trunk/gamemode/glon.lua
 */

// Type 2: Table
type TTable map[interface{}]interface{}

// Type 3: Array
type TArray []interface{}

// Type 4 & 5: false, true respectively
type TBool bool

// Type 6: Number 
type TNumber float64

// Type 7: String
type TString string

// Type 8: Vector
// Spec: http://maurits.tv/data/garrysmod/wiki/wiki.garrysmod.com/index3d06.html
type TVector struct {
	X, Y, Z float64
}

// Type 9: Angle
// // Spec: http://maurits.tv/data/garrysmod/wiki/wiki.garrysmod.com/index5012.html
type TAngle struct {
	P, Y, R float64
}

// Type 10: Entity
type TEntity interface{}

// Type 11: Player
type TPlayer interface{}

// Type 12: CEffectData
type TCEffectData interface{}

// Type 13: ConVar 
type TConVar interface{}

// Type 15: Color
type TColor struct {
	R, G, B, A float64
}

// Type 255: Reference
type TReference *interface{}
