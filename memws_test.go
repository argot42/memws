package memws

import "testing"

type SeekCase struct {
	NPos    int64
	Whence  int
	EPos    int64
	Success bool
}

func TestWrite(t *testing.T) {
	in := &WriteSeek{}
	out := []string{
		"漢字",
		"漢字foo",
		"漢字fbar",
	}

	in.Write([]byte(out[0]))
	if out[0] != in.String() {
		t.Errorf("%s should be %s", in, out[0])
	}
	in.Write([]byte("foo"))
	if out[1] != in.String() {
		t.Errorf("%s should be %s", in, out[1])
	}
	in.pos = 7
	in.Write([]byte("bar"))
	if out[2] != in.String() {
		t.Errorf("%s should be %s", in, out[2])
	}
}

func TestSeek(t *testing.T) {
	testCases := []SeekCase{
		SeekCase{
			-10,
			0,
			0,
			false,
		},
		SeekCase{
			10,
			0,
			10,
			true,
		},
		SeekCase{
			5,
			1,
			15,
			true,
		},
		SeekCase{
			-4,
			1,
			11,
			true,
		},
		SeekCase{
			10,
			2,
			21,
			true,
		},
		SeekCase{
			0,
			-1,
			0,
			false,
		},
		SeekCase{
			0,
			100,
			0,
			false,
		},
	}

	w := &WriteSeek{}

	for i, tc := range testCases {
		_, err := w.Seek(tc.NPos, tc.Whence)
		if err != nil {
			if tc.Success {
				t.Errorf("tc n° %d shouldn't have failed", i)
				break
			}
			return
		}

		if !tc.Success {
			t.Errorf("tc n° %d should have failed", i)
			break
		}

		if w.pos != tc.EPos {
			t.Errorf("tc n° %d position is %d but should be %d", i, w.pos, tc.EPos)
		}
	}
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
