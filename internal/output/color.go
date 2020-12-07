package output

import (
	"github.com/fatih/color"
)

var whiteBold *color.Color

func init() {
	whiteBold = color.New(color.FgWhite).Add(color.Bold)
}
