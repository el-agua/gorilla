package gorilla

import (
	"fmt"
	"io"
)

type MultiCompressor struct {
	compressors []Compressor
}

func NewMultiCompressor() (c *MultiCompressor, err error) {
	c = &MultiCompressor{
	}
	c.compressors = make([]Compressor, 0)
	return c, nil
}

func (c *MultiCompressor) Compress(i uint32, t uint32, v float64) error {
	if i >= uint32(len(c.compressors)) {
		return fmt.Errorf("index out of range: %d", i)
	}
	return c.compressors[i].Compress(t, v)
}

func (c *MultiCompressor) AddCompressor(w io.Writer, header uint32) error {
	compressor, _, err := NewCompressor(w, header)
	if err != nil {
		return err
	}
	c.compressors = append(c.compressors, *compressor)
	return nil
}