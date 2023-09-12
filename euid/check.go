package euid

func shiftRight7(v [2]uint64) [2]uint64 {
	var mask = uint64(0x7f)
	var a1 = v[0] & mask
	var a = v[0] >> 7
	var b = (v[1] >> 7) | (a1 << 57)
	return [2]uint64{a, b}
}

func addU128(a [2]uint64, b [2]uint64) [2]uint64 {
	var a1 = a[0]
	var a2 = a[1]
	var b1 = b[0]
	var b2 = b[1]

	var sum1 = uint64(0)
	var sum2 = uint64(0)
	var carry1 = uint64(0)
	var carry2 = uint64(1)

	for carry1 != 0 || carry2 != 0 {
		sum1 = a1 ^ b1
		sum2 = a2 ^ b2
		var a2b2 = a2 & b2
		carry2 = a2b2 << 1
		carry1 = ((a1 & b1) << 1) | (a2b2 >> 63)
		a1 = sum1
		a2 = sum2
		b1 = carry1
		b2 = carry2
	}
	return [2]uint64{sum1, sum2}
}

func subU128(a [2]uint64, b [2]uint64) [2]uint64 {
	var r = addU128([2]uint64{^b[0], ^b[1]}, [2]uint64{0, 1})
	return addU128(a, r)
}

func isGtP(v [2]uint64, p uint64) bool {
	if v[0] != 0 {
		return true
	} else {
		return v[1] > p
	}
}

func m7(euid EUID) u5 {
	var p = uint64(0x7f)
	var i = addU128([2]uint64{0, euid.lo & p}, shiftRight7([2]uint64{euid.hi, euid.lo}))
	for isGtP(i, p) {
		i = addU128([2]uint64{0, i[1] & p}, shiftRight7(i))
	}
	if i[0] == 0 && i[1] == p {
		return 0
	} else {
		return u5(i[1])
	}
}
