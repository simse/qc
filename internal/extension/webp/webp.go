package webp

import (
	"io"

	"golang.org/x/image/webp"

	"github.com/simse/qc/internal/format"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension: "webp",
		Aliases:   []string{},
		HumanName: "webp",
		Library:   "native",
		//Encoder:          Encode,
		Decoder:          Decode,
		EncoderAvailable: false,
		DecoderAvailable: true,
	}
}

// Decode converts the PNG image to a generic image object for encoding
func Decode(reader io.Reader) format.DecodeOutput {
	imageObject, imageReadError := webp.Decode(reader)
	if imageReadError != nil {
		panic(imageReadError)
	}

	return format.DecodeOutput{
		Image: imageObject,
	}
}

/*
// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject format.DecodeOutput) format.EncodeOutput {
	return format.EncodeOutput{}
}
*/
