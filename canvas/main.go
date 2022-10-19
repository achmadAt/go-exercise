package main

import (
	"image/png"
	"os"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

var fontFamily *canvas.FontFamily

func main() {
	fontFamily = canvas.NewFontFamily("times")
	if err := fontFamily.LoadLocalFont("NimbusRoman-Regular", canvas.FontRegular); err != nil {
		panic(err)
	}

	c := canvas.New(200, 230)
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(c.W, c.H))
	draw(ctx)
	renderers.Write("test2.png", c, canvas.DPMM(5.0))
}

var lorem = []string{
	`a cat inside garden flowers`,
}

var y = 205.0

func drawText(c *canvas.Context, x float64, text *canvas.Text) {
	h := text.Bounds().H
	c.DrawText(x, y, text)
	y -= h + 10.0
}

func draw(c *canvas.Context) {
	c.SetFillColor(canvas.Black)

	headerFace := fontFamily.Face(28.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	textFace := fontFamily.Face(12.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)

	drawText(c, 30.0, canvas.NewTextBox(headerFace, "Document Example", 0.0, 0.0, canvas.Left, canvas.Top, 0.0, 0.0))
	drawText(c, 30.0, canvas.NewTextBox(textFace, lorem[0], 140.0, 0.0, canvas.Justify, canvas.Top, 5.0, 0.0))

	lenna, err := os.Open("./some.png")
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(lenna)
	if err != nil {
		panic(err)
	}
	imgDPMM := 15.0
	imgWidth := float64(img.Bounds().Max.X) / imgDPMM
	imgHeight := float64(img.Bounds().Max.Y) / imgDPMM
	c.DrawImage(170.0-imgWidth, y-imgHeight, img, canvas.DPMM(imgDPMM))

	for text := range lorem {
		drawText(c, 30.0, canvas.NewTextBox(textFace, lorem[text], 140.0-imgWidth-10.0, 0.0, canvas.Justify, canvas.Top, 5.0, 0.0))
	}
}
