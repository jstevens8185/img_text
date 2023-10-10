// textoverlay.go

package img_text

import (
	"github.com/fogleman/gg"
)

// AddTextOverlay adds text on top of an existing image and saves it.
func AddTextOverlay(inputImagePath, outputImagePath, message string) error {
	// Load the existing image
	img, err := gg.LoadImage(inputImagePath)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	dc := gg.NewContext(width, height)

	// Set a background color (optional)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Draw the existing image
	dc.DrawImage(img, 0, 0)

	// Set text color and size using the default font
	dc.SetRGB(0.5, 0, 0)
	if err := dc.LoadFontFace("Roboto-Bold.ttf", 72); err != nil {
		return err
	}

	// Write text on top of the image
	dc.DrawStringAnchored(message, float64(width)/2, float64(height)/2, 0.5, 0.5)

	// Save the modified image with text
	if err := dc.SavePNG(outputImagePath); err != nil {
		return err
	}

	return nil
}
