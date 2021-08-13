package jpg

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/simse/qc/internal/format"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension:        "jpg",
		Aliases:          []string{"jpeg"},
		HumanName:        "jpg",
		Library:          "native",
		Encoder:          Encode,
		Decoder:          Decode,
		EncoderAvailable: true,
		DecoderAvailable: true,
	}
}

// Decode converts the JPG image to a generic image object for encoding
func Decode(reader io.Reader) (interface{}, error) {
	imageObject, decodeError := jpeg.Decode(reader)

	if decodeError != nil {
		return nil, decodeError
	}

	return imageObject, nil
}

// Encode converts a generic image object to a JPG file
func Encode(writer io.Writer, decodeInfo interface{}) (interface{}, error) {
	image := decodeInfo.(image.Image)

	convertError := jpeg.Encode(writer, image, &jpeg.Options{
		Quality: 90,
	})

	return format.EncodeOutput{
		Status: convertError == nil,
	}, nil
}
