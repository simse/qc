package png

import (
	"image"
	"io"

	fastpng "github.com/amarburg/go-fast-png"
	"github.com/simse/qc/internal/format"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension:        "png",
		Aliases:          []string{},
		HumanName:        "png",
		Library:          "fastpng",
		Encoder:          Encode,
		Decoder:          Decode,
		EncoderAvailable: true,
		DecoderAvailable: true,
	}
}

// Decode converts the PNG image to a generic image object for encoding
func Decode(reader io.Reader) (interface{}, error) {
	imageObject, _ := fastpng.Decode(reader)

	return imageObject, nil
}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject interface{}) (interface{}, error) {
	image := decodeObject.(image.Image)
	encoder := fastpng.Encoder{
		CompressionLevel: fastpng.BestSpeed,
	}

	convertError := encoder.Encode(writer, image)

	return format.EncodeOutput{
		Status: convertError == nil,
	}, nil
}
