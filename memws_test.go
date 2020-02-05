package memws

import "testing"

func TestWrite(t *testing.T) {
}

func TestSeek(t *testing.T) {
}

func TestString(t *testing.T) {
	in := &WriteSeek{
		[]byte("漢字"),
		0,
	}
	out := "漢字"

	if out != in.String() {
		t.Fatalf("string mismatch %s should be %s", in, out)
	}
}
