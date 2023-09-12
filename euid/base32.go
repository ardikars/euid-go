package euid

import (
	"errors"
)

type u5 uint8

const mcp = u5(255)

var encodingSymbols = [32]rune{
	'0', '1', '2', '3', '4', '5', '6', '7', //
	'8', '9', 'A', 'B', 'C', 'D', 'E', 'F', //
	'G', 'H', 'J', 'K', 'M', 'N', 'P', 'Q', //
	'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z', //
}

var decodingSymbols = [123]u5{
	mcp, mcp, mcp, mcp, mcp, mcp, mcp, mcp, // 0
	mcp, mcp, mcp, mcp, mcp, mcp, mcp, mcp, // 8
	mcp, mcp, mcp, mcp, mcp, mcp, mcp, mcp, // 16
	mcp, mcp, mcp, mcp, mcp, mcp, mcp, mcp, // 24
	mcp, mcp, mcp, mcp, mcp, mcp, mcp, mcp, // 32
	mcp, mcp, mcp, mcp, mcp, mcp, mcp, mcp, // 40
	0, 1, 2, 3, 4, 5, 6, 7, // 48
	8, 9, mcp, mcp, mcp, mcp, mcp, mcp, // 56
	mcp, 10, 11, 12, 13, 14, 15, 16, // 64
	17, 1, 18, 19, 1, 20, 21, 0, // 72
	22, 23, 24, 25, 26, mcp, 27, 28, // 80
	29, 30, 31, mcp, mcp, mcp, mcp, mcp, // 88
	mcp, 10, 11, 12, 13, 14, 15, 16, // 96
	17, 1, 18, 19, 1, 20, 21, 0, // 104
	22, 23, 24, 25, 26, mcp, 27, 28, // 112
	29, 30, 31, // 120
}

func toU5Slice(hi uint64, lo uint64) [27]u5 {
	var dst [27]u5
	dst[0] = u5((hi >> 59) & 0x1f)
	dst[1] = u5((hi >> 54) & 0x1f)
	dst[2] = u5((hi >> 49) & 0x1f)
	dst[3] = u5((hi >> 44) & 0x1f)
	dst[4] = u5((hi >> 39) & 0x1f)
	dst[5] = u5((hi >> 34) & 0x1f)
	dst[6] = u5((hi >> 29) & 0x1f)
	dst[7] = u5((hi >> 24) & 0x1f)
	dst[8] = u5((hi >> 19) & 0x1f)
	dst[9] = u5((hi >> 14) & 0x1f)
	dst[10] = u5((hi >> 9) & 0x1f)
	dst[11] = u5((hi >> 4) & 0x1f)
	dst[12] = u5(((hi & 0xf) << 1) | ((lo >> 63) & 0x1))
	//
	dst[13] = u5((lo >> 58) & 0x1f)
	dst[14] = u5((lo >> 53) & 0x1f)
	dst[15] = u5((lo >> 48) & 0x1f)
	dst[16] = u5((lo >> 43) & 0x1f)
	dst[17] = u5((lo >> 38) & 0x1f)
	dst[18] = u5((lo >> 33) & 0x1f)
	dst[19] = u5((lo >> 28) & 0x1f)
	dst[20] = u5((lo >> 23) & 0x1f)
	dst[21] = u5((lo >> 18) & 0x1f)
	dst[22] = u5((lo >> 13) & 0x1f)
	dst[23] = u5((lo >> 8) & 0x1f)
	dst[24] = u5((lo >> 3) & 0x1f)
	dst[25] = u5((lo & 0x7) << 2)
	return dst
}

func toU64Slice(slice [27]u5) (uint64, uint64) {
	var hi = (uint64(slice[0]) << 59) | (uint64(slice[1]) << 54) | (uint64(slice[2]) << 49) | (uint64(slice[3]) << 44) | (uint64(slice[4]) << 39) | (uint64(slice[5]) << 34) | (uint64(slice[6]) << 29) | (uint64(slice[7]) << 24) | (uint64(slice[8]) << 19) | (uint64(slice[9]) << 14) | (uint64(slice[10]) << 9) | (uint64(slice[11]) << 4) | (uint64(slice[12]>>1) & 0xf)
	var lo = (uint64(slice[13]) << 58) | ((uint64(slice[12]) & 0x1) << 63) | (uint64(slice[14]) << 53) | (uint64(slice[15]) << 48) | (uint64(slice[16]) << 43) | (uint64(slice[17]) << 38) | (uint64(slice[18]) << 33) | (uint64(slice[19]) << 28) | (uint64(slice[20]) << 23) | (uint64(slice[21]) << 18) | (uint64(slice[22]) << 13) | (uint64(slice[23]) << 8) | (uint64(slice[24]) << 3) | (uint64(slice[25]) >> 2)
	return hi, lo
}

func encode(euid EUID, checkmod bool) string {
	var slice = toU5Slice(euid.hi, euid.lo)
	var dst [27]rune
	dst[0] = encodingSymbols[slice[0]]
	dst[1] = encodingSymbols[slice[1]]
	dst[2] = encodingSymbols[slice[2]]
	dst[3] = encodingSymbols[slice[3]]
	dst[4] = encodingSymbols[slice[4]]
	dst[5] = encodingSymbols[slice[5]]
	dst[6] = encodingSymbols[slice[6]]
	dst[7] = encodingSymbols[slice[7]]
	dst[8] = encodingSymbols[slice[8]]
	dst[9] = encodingSymbols[slice[9]]
	dst[10] = encodingSymbols[slice[10]]
	dst[11] = encodingSymbols[slice[11]]
	dst[12] = encodingSymbols[slice[12]]
	dst[13] = encodingSymbols[slice[13]]
	dst[14] = encodingSymbols[slice[14]]
	dst[15] = encodingSymbols[slice[15]]
	dst[16] = encodingSymbols[slice[16]]
	dst[17] = encodingSymbols[slice[17]]
	dst[18] = encodingSymbols[slice[18]]
	dst[19] = encodingSymbols[slice[19]]
	dst[20] = encodingSymbols[slice[20]]
	dst[21] = encodingSymbols[slice[21]]
	dst[22] = encodingSymbols[slice[22]]
	dst[23] = encodingSymbols[slice[23]]
	dst[24] = encodingSymbols[slice[24]]
	var check u5
	if checkmod {
		check = m7(euid)
	} else {
		check = u5(0x7f)
	}
	dst[25] = encodingSymbols[slice[25]|(check>>5)]
	dst[26] = encodingSymbols[check&0x1f]
	str := string(dst[:])
	return str
}

func decode(encoded string) (EUID, error) {
	if len(encoded) != 27 {
		return EUID{0, 0}, errors.New("invalid length")
	}
	var slice [27]u5
	for index, char := range encoded {
		var codePoint = int(char)
		if codePoint > len(decodingSymbols) {
			return EUID{0, 0}, errors.New("invalid character")
		}
		slice[index] = u5(decodingSymbols[codePoint])
		if slice[index] == 255 {
			return EUID{0, 0}, errors.New("invalid character")
		}
	}
	var r = slice[25] & 0x3
	slice[25] &= 0x1c
	var hi, lo = toU64Slice(slice)
	var check = (r << 5) | slice[26]
	if check == 0x7f {
		return EUID{hi, lo}, nil
	} else {
		var euid = EUID{hi, lo}
		var w = m7(euid)
		if check != w {
			return EUID{0, 0}, errors.New("invalid checkmod")
		} else {
			return euid, nil
		}
	}
}
