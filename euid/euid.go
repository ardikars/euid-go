package euid

import "errors"

const timestampBitmask = 0x1fffffffffff
const extLenBitmask = 0xf
const extDataBitmask = 0x7fff

type EUID struct {
	hi uint64
	lo uint64
}

func Create() (EUID, error) {
	var ts = currentTimestamp()
	var euid, err = createWithTimestamp(ts)
	if err != nil {
		return EUID{0, 0}, err
	}
	return euid, nil
}

func CreateWithExtension(ext uint16) (EUID, error) {
	if ext > extDataBitmask {
		return EUID{0, 0}, errors.New("invalid extension")
	} else {
		var ts = currentTimestamp()
		var euid, err = createWithTimestampAndExtension(ts, ext)
		if err != nil {
			return EUID{0, 0}, err
		}
		return euid, nil
	}
}

func FromString(str string) (EUID, error) {
	var euid, err = decode(str)
	if err != nil {
		return EUID{0, 0}, err
	}
	return euid, nil
}

func FromBytes(bytes []byte) (EUID, error) {
	if bytes == nil || len(bytes) != 16 {
		return EUID{0, 0}, errors.New("invalid length")
	}
	var hi = (uint64(bytes[0]) << 56) | (uint64(bytes[1]) << 48) | (uint64(bytes[2]) << 40) | (uint64(bytes[3]) << 32) | (uint64(bytes[4]) << 24) | (uint64(bytes[5]) << 16) | (uint64(bytes[6]) << 8) | uint64(bytes[7])
	var lo = (uint64(bytes[8]) << 56) | (uint64(bytes[9]) << 48) | (uint64(bytes[10]) << 40) | (uint64(bytes[11]) << 32) | (uint64(bytes[12]) << 24) | (uint64(bytes[13]) << 16) | (uint64(bytes[14]) << 8) | uint64(bytes[15])
	return EUID{hi, lo}, nil
}

func (euid EUID) Extension() (uint16, error) {
	var extLen = euid.hi & extLenBitmask
	if extLen == 0 {
		return 0, errors.New("no extension")
	}
	var bitmask = uint64((1 << extLen) - 1)
	return uint16((euid.hi >> 4) & bitmask), nil
}

func (euid EUID) Timestamp() uint64 {
	return (euid.hi >> 19) & timestampBitmask
}

func (euid EUID) Next() (EUID, error) {
	var timestamp = currentTimestamp()
	if timestamp == euid.Timestamp() {
		var rHi = euid.hi >> 32
		if rHi == 0xffffffff {
			return EUID{0, 0}, errors.New("timestamp overflow")
		} else {
			return EUID{euid.hi, ((rHi + 1) << 32) | uint64(random32())}, nil
		}
	} else {
		var ext, err = euid.Extension()
		if err != nil {
			return createWithTimestamp(timestamp)
		} else {
			return createWithTimestampAndExtension(timestamp, ext)
		}
	}
}

func (euid EUID) ToBytes() []byte {
	var bytes = make([]byte, 16)
	bytes[0] = byte((euid.hi >> 56) & 0xff)
	bytes[1] = byte((euid.hi >> 48) & 0xff)
	bytes[2] = byte((euid.hi >> 40) & 0xff)
	bytes[3] = byte((euid.hi >> 32) & 0xff)
	bytes[4] = byte((euid.hi >> 24) & 0xff)
	bytes[5] = byte((euid.hi >> 16) & 0xff)
	bytes[6] = byte((euid.hi >> 8) & 0xff)
	bytes[7] = byte(euid.hi & 0xff)
	bytes[8] = byte((euid.lo >> 56) & 0xff)
	bytes[9] = byte((euid.lo >> 48) & 0xff)
	bytes[10] = byte((euid.lo >> 40) & 0xff)
	bytes[11] = byte((euid.lo >> 32) & 0xff)
	bytes[12] = byte((euid.lo >> 24) & 0xff)
	bytes[13] = byte((euid.lo >> 16) & 0xff)
	bytes[14] = byte((euid.lo >> 8) & 0xff)
	bytes[15] = byte(euid.lo & 0xff)
	return bytes
}

func (euid EUID) Encode(checkmod bool) string {
	return encode(euid, checkmod)
}

func (euid EUID) ToString() string {
	return euid.Encode(true)
}

func createWithTimestamp(timestamp uint64) (EUID, error) {
	if timestamp > timestampBitmask {
		return EUID{0, 0}, errors.New("timestamp overflow")
	} else {
		var r = random2U64()
		return EUID{(timestamp << 19) | ((r[0] & 0x7fff) << 4), r[1]}, nil
	}
}

func createWithTimestampAndExtension(timestamp uint64, extension uint16) (EUID, error) {
	if timestamp > timestampBitmask {
		return EUID{0, 0}, errors.New("timestamp overflow")
	} else {
		var extData = uint64(extension)
		var extLen = getExtBitLen(extension)
		var r = random2U64()
		var ramainingRand = r[0] & ((1 << (15 - extLen)) - 1)
		var hi = (timestamp << 19) | (ramainingRand << (4 + extLen)) | (extData << 4) | extLen
		return EUID{hi, r[1]}, nil
	}
}

func getExtBitLen(ext uint16) uint64 {
	var x = ext & 0x7fff
	var n = uint64(0)
	if x <= 0x00ff {
		n += 8
		x <<= 8
	}
	if x <= 0x0fff {
		n += 4
		x <<= 4
	}
	if x <= 0x3fff {
		n += 2
		x <<= 2
	}
	if x <= 0x7fff {
		n += 1
	}
	return uint64(16 - n)
}
