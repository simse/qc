package images

import "github.com/davidbyttow/govips/v2/vips"

// Check checks whether libvips is installed and ready
func Check() bool {
	return true
}

// Setup starts libvips
func Setup() bool {
	// Start lipvips
	startupConfig := vips.Config{
		ConcurrencyLevel: 8,
	}
	vips.Startup(&startupConfig)
	// defer vips.Shutdown()

	return true
}
