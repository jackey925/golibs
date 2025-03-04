package golibs

import (
	"bytes"
	"fmt"
)

type StringBuilder struct {
	buf bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{buf: bytes.Buffer{}}
}

func (s *StringBuilder) Append(obj interface{}) *StringBuilder {
	s.buf.WriteString(fmt.Sprintf("%v", obj))
	return s
}

func (s *StringBuilder) ToString() string {
	return s.buf.String()
}
func (s *StringBuilder) AppendForString(obj string) *StringBuilder {
	s.buf.WriteString(obj)
	return s
}


func (s *StringBuilder) AppendForBytes(obj []byte) *StringBuilder {
	s.buf.Write(obj)
	return s
}