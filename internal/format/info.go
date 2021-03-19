package format

import "io"

// Info stores information about a format
type Info struct {
	Library          string
	Extension        string
	Aliases          []string
	HumanName        string
	Decoder          func(io.Reader) (interface{}, error)
	Encoder          func(io.Writer, interface{}) (interface{}, error)
	DecoderAvailable bool
	EncoderAvailable bool
}
