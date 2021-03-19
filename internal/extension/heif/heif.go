package heif

import (
	"io"

	goheif "github.com/simse/go-heif"
	"github.com/simse/qc/internal/format"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension: "heif",
		Aliases:   []string{"heic"},
		HumanName: "heif",
		Library:   "native",
		//Encoder:          Encode,
		Decoder:          Decode,
		EncoderAvailable: false,
		DecoderAvailable: true,
	}
}

// Decode converts the JPG image to a generic image object for encoding
func Decode(reader io.Reader) (interface{}, error) {
	imageObject, _ := goheif.Decode(reader)
	// panic(err)

	return imageObject, nil
}

/*
// Encode converts a generic image object to a JPG file
func Encode(writer io.Writer, decodeInfo interface{}) (interface{}, error) {
	image := decodeInfo.(*image.Image)



	convertError := goheif.Encode()

	return format.EncodeOutput{
		Status: convertError == nil,
	}, nil
}
*/
