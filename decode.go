package austere

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
)

// Decoder read data from reader r and decode
type Decoder struct {
	r io.Reader
}

// NewDecoderWithBuffer returns a Decoder that read
// from r with buffer
func NewDecoderWithBuffer(r io.Reader) *Decoder {
	return &Decoder{
		r: bufio.NewReader(r),
	}
}

// NewDecoderWithBufferSize returns a Decoder that
// read from r with buffer, assign buffer size with
// size
func NewDecoderWithBufferSize(r io.Reader, size int) *Decoder {
	return &Decoder{
		r: bufio.NewReaderSize(r, size),
	}
}

// Decode
func (d *Decoder) ReadAndDecode() ([]byte, error) {
	if d.r == nil {
		return nil, errors.New("Decoder has invalid reader r")
	}

	var length uint64
	err := binary.Read(d.r, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	payload := make([]byte, length)
	err = binary.Read(d.r, binary.BigEndian, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
