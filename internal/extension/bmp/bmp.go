package bmp

import (
	"io"

	"github.com/simse/qc/internal/format"
	"golang.org/x/image/bmp"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension: "bmp",
		Aliases:   []string{},
		HumanName: "bmp",
		Library:   "native",
		Encoder:   Encode,
		// Decoder:   Decode,
		EncoderAvailable: true,
		DecoderAvailable: true,
	}
}

// Decode converts the BMP image to a generic image object for encoding
func Decode(reader io.Reader) format.DecodeOutput {
	imageObject, _ := bmp.Decode(reader)

	return format.DecodeOutput{
		Image: imageObject,
	}
}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject format.DecodeOutput) format.EncodeOutput {
	writeError := bmp.Encode(writer, decodeObject.Image)

	return format.EncodeOutput{
		Status: writeError == nil,
	}
}
