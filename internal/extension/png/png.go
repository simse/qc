package png

import (
	"io"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/simse/qc/internal/format"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension:        "png",
		Aliases:          []string{},
		HumanName:        "png",
		Library:          "libvips",
		Encoder:          Encode,
		Decoder:          Decode,
		EncoderAvailable: true,
		DecoderAvailable: true,
	}
}

// Decode converts the PNG image to a generic image object for encoding
func Decode(reader io.Reader) (interface{}, error) {
	imageObject, decodeError := vips.NewImageFromReader(reader)

	if decodeError != nil {
		return nil, decodeError
	}

	return imageObject, nil
}

// Encode converts a generic image object to a PNG file
func Encode(writer io.Writer, decodeObject interface{}) (interface{}, error) {
	exportParameters := vips.PngExportParams{
		StripMetadata: true,
		Compression:   5,
		Interlace:     false,
	}

	pngImage, _, convertError := decodeObject.(*vips.ImageRef).ExportPng(&exportParameters)

	writer.Write(pngImage)

	return format.EncodeOutput{
		Status: convertError == nil,
	}, nil
}
