package main

import (
    "fmt"
    "image"
	"os"
    "time"

    _ "image/png"

    "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	
	"golang.org/x/image/font/basicfont"
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
    type GameState int
	const (
	    GameStateTitleScreen GameState = iota
		GameStatePlaying
		GameStateGameOver
	)
	
	gameState := GameStateTitleScreen
	
	type GameScreen int
	const (
	    GameScreenVillage GameScreen = iota
		GameScreenForest
	)
	
	gameScreen := GameScreenVillage
	
	// Load background images
	
    backgroundTitlePic, err := loadPicture("assets/backgrounds/title.png")
	if err != nil {
		panic(err)
	}
    backgroundTitleSprite := pixel.NewSprite(backgroundTitlePic, backgroundTitlePic.Bounds())

    backgroundVillagePic, err := loadPicture("assets/backgrounds/village.png")
	if err != nil {
		panic(err)
	}
	backgroundVillageSprite := pixel.NewSprite(backgroundVillagePic, backgroundVillagePic.Bounds())

    cfg := pixelgl.WindowConfig{
		Title:  "Adventure Game",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

    basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	titleInfoTxt := text.New(pixel.V(800,50), basicAtlas)
	fmt.Fprintln(titleInfoTxt, "Created by James Stocks")
	titlePromptTxt := text.New(pixel.V(400,300), basicAtlas)
	fmt.Fprintln(titlePromptTxt, "Press Enter to start your adventure")

    fps60 := time.Tick(time.Second / 60)

	for !win.Closed() {
	    // Update the game
		switch gameState {
		case GameStateTitleScreen:
			if win.JustPressed(pixelgl.KeyEnter) {
		        gameState = GameStatePlaying
		    }
		}
		
	    // Draw the screen
	    win.Clear(pixel.RGB(0, 0, 0))
		
		switch gameState {
		case GameStateTitleScreen:
		    backgroundTitleSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		    titleInfoTxt.Draw(win, pixel.IM)
			titlePromptTxt.Draw(win, pixel.IM)
        case GameStatePlaying:
		    switch gameScreen {
			case GameScreenVillage:
			    backgroundVillageSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
			}
		}
		
		win.Update()

		<-fps60
	}
}

func main() {
	pixelgl.Run(run)
}

