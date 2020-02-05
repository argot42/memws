package memws

import (
	"errors"
	"io"
)

type WriteSeek struct {
	buf []byte
	pos int64
}

func (ws *WriteSeek) Write(p []byte) (int, error) {
	c := ws.pos + int64(len(p))
	if c > int64(cap(ws.buf)) {
		newBuf := make([]byte, len(ws.buf), c+int64(len(p)))
		copy(newBuf, ws.buf)
		ws.buf = newBuf
	}
	if c > int64(len(ws.buf)) {
		ws.buf = ws.buf[:c]
	}
	copy(ws.buf[ws.pos:], p)
	ws.pos += int64(len(p))
	return len(p), nil
}

func (ws *WriteSeek) Seek(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = offset + ws.pos
	case io.SeekEnd:
		abs = offset + int64(len(ws.buf))
	default:
		return 0, errors.New("writeseek.WriteSeek.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("writeseek.WriteSeek.Seek: negative pos")
	}
	ws.pos = abs
	return abs, nil
}

func (ws WriteSeek) String() string {
	return string(ws.buf)
}
