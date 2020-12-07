package webp

import (
	"io"

	"github.com/chai2010/webp"
	"github.com/simse/qc/internal/format"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension: "webp",
		Aliases:   []string{},
		HumanName: "webp",
		Encoder:   Encode,
		// Decoder:   Decode,
	}
}

// Decode converts the PNG image to a generic image object for encoding
func Decode() {

}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject format.DecodeOutput) format.EncodeOutput {
	writeError := webp.Encode(writer, decodeObject.Image, nil)

	return format.EncodeOutput{
		Status: writeError == nil,
	}
}
