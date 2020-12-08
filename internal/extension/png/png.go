package png

import (
	"io"

	fastpng "github.com/amarburg/go-fast-png"
	"github.com/simse/qc/internal/format"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension: "png",
		Aliases:   []string{},
		HumanName: "png",
		Library:   "native",
		Encoder:   Encode,
		// Decoder:   Decode,
		EncoderAvailable: true,
		DecoderAvailable: true,
	}
}

// Decode converts the PNG image to a generic image object for encoding
func Decode() {

}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject format.DecodeOutput) format.EncodeOutput {
	pngEncoder := fastpng.Encoder{
		CompressionLevel: fastpng.BestSpeed,
	}

	writeError := pngEncoder.Encode(writer, decodeObject.Image)

	return format.EncodeOutput{
		Status: writeError == nil,
	}
}
