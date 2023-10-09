package img_text

import (
	"github.com/fogleman/gg"
)

func AddTextToImage(outputImagePath, text string) error {
	const W = 500
	const H = 300
	const fontSize = 72

	dc := gg.NewContext(W, H)

	// Clear the background
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Set the text color
	dc.SetRGB(.5, 0, 0)

	// Load the font
	if err := dc.LoadFontFace("path/to/your/font.ttf", fontSize); err != nil {
		return err
	}

	// Calculate text width and height
	textWidth, textHeight := dc.MeasureString(text)

	// Calculate text position to center it
	x := (W - textWidth) / 2
	y := (H - textHeight) / 2

	// Draw the text on the image
	dc.DrawStringAnchored(text, x+textWidth/2, y+textHeight/2, 0.5, 0.5)
	dc.Stroke()

	// Save the resulting image
	if err := dc.SavePNG(outputImagePath); err != nil {
		return err
	}

	return nil
}
