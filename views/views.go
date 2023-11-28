package views

import (
	"carro/models"
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	imDraw "github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	colorNames "golang.org/x/image/colornames"
)

var (
	background *pixel.Sprite
	bgPicture  pixel.Picture
)

// loadBackground carga la imagen de fondo.
func loadBackground() {
	file, err := os.Open("assets/background.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	bgPicture = pixel.PictureDataFromImage(img)
	background = pixel.NewSprite(bgPicture, bgPicture.Bounds())
}
func drawCars(imd *imDraw.IMDraw, cars []models.Car) {
	for _, car := range cars {
		imd.Color = colorNames.Red

		carWidth := 20.0
		carHeight := 20.0

		topLeft := car.Position.Add(pixel.V(-carWidth/2, -carHeight/2))
		topRight := car.Position.Add(pixel.V(carWidth/2, -carHeight/2))
		bottomLeft := car.Position.Add(pixel.V(-carWidth/2, carHeight/2))
		bottomRight := car.Position.Add(pixel.V(carWidth/2, carHeight/2))

		imd.Push(topLeft, topRight, bottomRight, bottomLeft)
		imd.Polygon(0)
	}
}

func DrawParkingLot(win *pixelgl.Window, cars []models.Car) {
	if background == nil {
		loadBackground()
	}

	background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	imd := imDraw.New(nil)
	imd.Color = colorNames.White
	drawCars(imd, cars)

	imd.Draw(win)
}
