package core

import (
	"fmt"
	"io"
	"strconv"
	"sync"
	"text2binary/core/util"
)

// encoder Provides functions for encoding
// data into its binary value.
type encoder struct {
	readBuf  []byte
	mu sync.Mutex
}

// ConvertAndWrite uses reader and writer interfaces to stream encoded binary data.
//
//
// [src] Is the source data reader to be encoded in binary.
//
// [dst] Is the destination writer for the binary-encoded data.
//
// [bufLen] Is the length of the memory buffer for reading. This field can be omitted with a value <= 0. If omitted, a default buffer length of 8192 bytes will be used.
//
// [delim] Is the delimiter sequence used to separate the binary values of each byte. This field can be omitted with nil. (no delimiter)
func (c *encoder) ConvertAndWrite(src io.Reader, dst io.Writer, bufLen int64, delim []byte) (err error) {
	// Cheeky nil pointer dereference prevention strat ( ͡° ͜ʖ ͡°)
	switch {
	case src == nil:
		fallthrough
	case dst == nil:
		return fmt.Errorf("src and dst must not be nil")
	}

	// Lock the mutex to prevent separate routines from messing up our buffer
	c.mu.Lock()
	defer c.mu.Unlock()

	// Initialize the buffer
	if bufLen > 0 {
		c.readBuf = make([]byte, bufLen)
	} else {
		c.readBuf = make([]byte, 8192)
	}

	// n will be populated
	var n int
	for {
		// Read up to len(c.readBuf) bytes into the buffer
		if n, err = src.Read(c.readBuf); err != nil {
			break
		}

		// Write the binary-encoded data to the buffer
		if _, err = dst.Write(c.Encode(c.readBuf[:n]).Delim(delim).Bytes()); err != nil {
			break
		}
	}

	// Write a newline byte because no one
	// likes that ugly EOL '%' character in ZSH
	_, _ = dst.Write([]byte{'\n'})

	// If it's EOF, that's expected so just return a nil error.
	// If not, something is definitely wrong - so we return the actual error
	if err == io.EOF || err == nil {
		return nil
	} else {
		return err
	}
}

// Result allows for multiple return types of the same binary result.
type Result struct {
	val []int
	delim []byte
}

// NewEncoder initializes a new encoder which can be used to convert data to
// a binary format.
func NewEncoder() *encoder {
	return &encoder{
		readBuf:  nil,
	}
}

// Encode encodes [b []byte] into a slice of []int values, which can be
// returned as a number of other types.
func (c *encoder) Encode(b []byte) *Result {
	res := &Result{
		val: make([]int, 0),
	}
	for _, v := range b {
		// Convert the rune to an integer, get the binary representation,
		// and append it to the result slice.
		res.val = append(res.val, util.ToBinary(int(v)))
	}
	return res
}

// Delim sets the delimiter & returns a pointer to Result
func (r *Result) Delim(delim []byte) *Result {
	r.delim = delim
	return r
}

// String returns the value of *Result in a string format.
func (r *Result) String() (s string) {
	if r.delim != nil {
		for i, v := range r.val {
			if i == len(r.val)-1 {
				s += strconv.Itoa(v)
				break
			}
			s += fmt.Sprintf("%s%s", strconv.Itoa(v), r.delim)
		}
		return
	}
	for _, v := range r.val {
		s += strconv.Itoa(v)
	}
	return
}

// Bytes returns the value of *Result in a byte slice.
func (r *Result) Bytes() (b []byte) {
	if r.delim != nil {
		for i, v := range r.val {
			if i == len(r.val)-1 {
				b = strconv.AppendInt(b, int64(v), 10)
				break
			}
			b = append(strconv.AppendInt(b, int64(v), 10), r.delim...)
		}
		return
	}
	for _, v := range r.val {
		b = strconv.AppendInt(b, int64(v), 10)
	}
	return
}

