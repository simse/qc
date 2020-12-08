package format

import "io"

// Info stores information about a format
type Info struct {
	Library          string
	Extension        string
	Aliases          []string
	HumanName        string
	Decoder          func(io.Reader) DecodeOutput
	Encoder          func(io.Writer, DecodeOutput) EncodeOutput
	DecoderAvailable bool
	EncoderAvailable bool
}
