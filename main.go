package main

import (
    "image"
	"os"
    "time"

    _ "image/png"

    "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
    backgroundTitlePic, err := loadPicture("assets/backgrounds/title.png")
	if err != nil {
		panic(err)
	}

    backgroundTitleSprite := pixel.NewSprite(backgroundTitlePic, backgroundTitlePic.Bounds())

    cfg := pixelgl.WindowConfig{
		Title:  "Adventure Game",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

    fps60 := time.Tick(time.Second / 60)

	for !win.Closed() {
	    win.Clear(pixel.RGB(0, 0, 0))
		
		backgroundTitleSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		
		win.Update()

		<-fps60
	}
}

func main() {
	pixelgl.Run(run)
}

