package jpg

import (
	"image/jpeg"
	"io"

	"github.com/simse/qc/internal/format"
)

// Info returns an Info struct about this format
func Info() format.Info {
	return format.Info{
		Extension: "jpg",
		Aliases:   []string{"jpeg"},
		HumanName: "jpg",
		// Encoder:   Encode,
		Decoder: Decode,
	}
}

// Decode converts the JPG image to a generic image object for encoding
func Decode(reader io.Reader) format.DecodeOutput {
	//imageFile, _ := os.Open(path)
	//bufferedReader := bufio.NewReader(f)
	imageObject, _ := jpeg.Decode(reader)
	// panic(err)

	return format.DecodeOutput{
		Image: imageObject,
	}
}

// Encode converts a generic image object to a JPG file
func Encode() {

}
