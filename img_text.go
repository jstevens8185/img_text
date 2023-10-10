package img_text

import (
	"log"

	"github.com/fogleman/gg"
)

func AddTextToImage(inputImage string, outputFile string) {
	const W = 500
	const H = 300

	// Load the existing image
	img, err := gg.LoadImage(inputImage)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(W, H)

	// Set a background color (optional)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Draw the existing image
	dc.DrawImage(img, 0, 0)

	// Set text color and size
	dc.SetRGB(0.5, 0, 0)
	if err := dc.LoadFontFace("Roboto-Bold.ttf", 72); err != nil {
		log.Fatal(err)
	}

	// Write text on top of the image
	message := "Hello, world!"
	dc.DrawStringAnchored(message, W/2, H/2, 0.5, 0.5)

	// Save the modified image with text
	if err := dc.SavePNG(outputFile); err != nil {
		log.Fatal(err)
	}
}
