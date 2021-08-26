package jpg

import (
	"io"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/simse/qc/internal/format"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension:        "jpg",
		Aliases:          []string{"jpeg"},
		HumanName:        "jpg",
		Library:          "libvips",
		Encoder:          Encode,
		Decoder:          Decode,
		EncoderAvailable: true,
		DecoderAvailable: true,
	}
}

// Decode converts the JPG image to a generic image object for encoding
func Decode(reader io.Reader) (interface{}, error) {
	imageObject, decodeError := vips.NewImageFromReader(reader)

	if decodeError != nil {
		return nil, decodeError
	}

	return imageObject, nil
}

// Encode converts a generic image object to a JPG file
func Encode(writer io.Writer, decodeObject interface{}) (interface{}, error) {
	exportParameters := vips.JpegExportParams{
		StripMetadata: true,
		Quality:       7,
		Interlace:     false,
	}
	jpgImage, _, convertError := decodeObject.(*vips.ImageRef).ExportJpeg(&exportParameters)

	writer.Write(jpgImage)

	return format.EncodeOutput{
		Status: convertError == nil,
	}, nil
}
