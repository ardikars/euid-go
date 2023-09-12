package euid

import (
	"crypto/rand"
	"encoding/binary"
)

func random2U64() [2]uint64 {
	b := make([]byte, 16)
	rand.Read(b)
	hi := binary.BigEndian.Uint64(b[8:])
	lo := binary.BigEndian.Uint64(b[0:8])
	return [2]uint64{hi, lo}
}

func random32() uint32 {
	b := make([]byte, 4)
	rand.Read(b)
	return binary.BigEndian.Uint32(b[8:])
}
