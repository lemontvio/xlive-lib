package avlib

import (
	"context"
	"io"
)

type AVIO interface {
	Now() (string, uint32, float64, uint32, error)
	Next(string) (string, uint32, float64, uint32, error)
	SetSize(int)
	Size() int
	String() string
	Len() int
	ToM3u8(func(int, string) string) string

	WriteToStream(context.Context, io.Writer) error
	WriteToSegment(io.Writer, string) error

	Write([]byte, float64) error
	WriteRange([]byte, float64, int) error
	Write2([]byte, float64, int, string) error

	Do(int, int, func(string, bool))
	Close() error
}
