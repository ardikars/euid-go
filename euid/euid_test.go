package euid

import (
	"testing"
)

func TestCreate(t *testing.T) {
	var euid, err = Create()
	if err != nil {
		t.Fatalf("Err (%v)\n", err)
	} else {
		var _, err = euid.Extension()
		if err == nil {
			t.Fatal("must no extention")
		} else {
			var timestmap = euid.Timestamp()
			t.Logf("t: EUID(%v), Timestamp(%v)\n", euid.ToString(), timestmap)
		}
	}
}

func TestCreateWithExtension(t *testing.T) {
	var euid0, err0 = CreateWithExtension(0)
	if err0 != nil {
		t.Fatalf("Err (%v)\n", err0)
	} else {
		var ext0, err0a = euid0.Extension()
		if err0a != nil {
			t.Fatalf("Err (%v)\n", err0a)
		} else {
			var timestmap0 = euid0.Timestamp()
			if ext0 != 0 {
				t.Fatalf("Invalid ext (%v)\n", ext0)
			} else {
				t.Logf("t: EUID(%v), Timestamp(%v), Extansion(%v)\n", euid0.ToString(), timestmap0, ext0)
			}
		}
	}
	var euid1, err1 = CreateWithExtension(32767)
	if err1 != nil {
		t.Fatalf("Err (%v)\n", err1)
	} else {
		var ext1, err1a = euid1.Extension()
		if err1a != nil {
			t.Fatalf("Err (%v)\n", err1a)
		} else {
			var timestmap1 = euid1.Timestamp()
			if ext1 != 32767 {
				t.Fatalf("Invalid ext (%v)\n", ext1)
			} else {
				t.Logf("t: EUID(%v), Timestamp(%v), Extansion(%v)\n", euid1.ToString(), timestmap1, ext1)
			}
		}
	}
	var _, err2 = CreateWithExtension(32768)
	if err2 == nil {
		t.Fatal("Must be fail, invalid ext")
	} else {
		t.Logf("Passed (%v)", err2)
	}
}

func TestNext(t *testing.T) {
	var euid0, err0 = Create()
	if err0 == nil {
		var _, err0a = euid0.Next()
		if err0a == nil {
			t.Logf("Passed")
		} else {
			t.Fatalf("Err (%v)", err0a)
		}
	} else {
		t.Fatalf("Err (%v)", err0)
	}
	var euid1, err1 = CreateWithExtension(2)
	if err1 == nil {
		var euid1a, err1a = euid1.Next()
		if err1a == nil {
			var a, erra = euid1.Extension()
			var b, errb = euid1a.Extension()
			if a == b && erra == nil && errb == nil {
				t.Logf("Passed")
			}
		} else {
			t.Fatalf("Err (%v)", err1a)
		}
	} else {
		t.Fatalf("Err (%v)", err1)
	}
}

func TestToBytesAndFromBytes(t *testing.T) {
	var euid, err = Create()
	if err != nil {
		t.Fatalf("Err (%v)", err)
	} else {
		var bytes = euid.ToBytes()
		var fromBytes, err1 = FromBytes(bytes)
		if err1 != nil {
			t.Fatalf("Err1 (%v)", err)
		} else {
			if euid != fromBytes {
				t.Fatalf("Err ne (%v)", err)
			}
		}
	}
}
