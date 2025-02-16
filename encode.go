// austere is a mini TCP Application layer protocol,
// only have a body length field in protocol header.
//
// This package aims to solve TCP sticky
// packet. But it's not production ready since
// it don't have any data verification means.
//
// The checksum field in TCP header may not
// guarantee data integrity in all scenarios.
// Do not trust it too much.

package austere

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
)

// Encoder encode data and write to w
type Encoder struct {
	w io.Writer
}

// NewEncoderWithBuffer returns a Encoder that write
// to w with buffer
func NewEncoderWithBuffer(w io.Writer) *Encoder {
	return &Encoder{
		w: bufio.NewWriter(w),
	}
}

// NewEncoderWithBufferSize returns a Encoder that
// write to w with buffer, assign buffer size with
// size
func NewEncoderWithBufferSize(w io.Writer, size int) *Encoder {
	return &Encoder{
		w: bufio.NewWriterSize(w, size),
	}
}

// EncodeAndWrite encode b and write to e.w,
// Don't forget to flush after writing if you use
// Encoder with buffer.
func (e *Encoder) EncodeAndWrite(b []byte) error {
	if e.w == nil {
		return errors.New("Encoder has invalid writer w")
	}

	length := uint64(len(b))

	err := binary.Write(e.w, binary.BigEndian, length)
	if err != nil {
		return err
	}

	err = binary.Write(e.w, binary.BigEndian, b)
	if err != nil {
		return err
	}

	return nil
}

// if e.w is a *bufio.Writer, write buffer.
// if isn't, do nothing.
func (e *Encoder) Flush() error {
	w, ok := e.w.(*bufio.Writer)
	if ok {
		return w.Flush()
	}

	return nil
}
