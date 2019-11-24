package main

import (
	"fmt"
	"image"
	"os"
	"time"

	"image/color"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"

	"golang.org/x/image/font/basicfont"

	"github.com/james-stocks/simple-game/player"
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

	backgroundForestPic, err := loadPicture("assets/backgrounds/forest.png")
	if err != nil {
		panic(err)
	}
	backgroundForestSprite := pixel.NewSprite(backgroundForestPic, backgroundForestPic.Bounds())

	blackPicturePic, err := loadPicture("assets/sprites/black64x64.png")
	if err != nil {
		panic(err)
	}
	blackSprite := pixel.NewSprite(blackPicturePic, blackPicturePic.Bounds())

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
	titleInfoTxt := text.New(pixel.V(800, 50), basicAtlas)
	fmt.Fprintln(titleInfoTxt, "Created by James Stocks")
	titlePromptTxt := text.New(pixel.V(400, 300), basicAtlas)
	fmt.Fprintln(titlePromptTxt, "Press Enter to start your adventure")

	villagePromptTxt := text.New(pixel.V(800, 50), basicAtlas)
	fmt.Fprintln(villagePromptTxt, "-> Begin Quest")

	questStepsRemainingTxt := text.New(pixel.V(800, 700), basicAtlas)

	playerLevelTxt := text.New(pixel.V(20, 700), basicAtlas)

	fps60 := time.Tick(time.Second / 60)

	player := player.Player{1, 0}
	isInQuest := false
	questStepsRemaining := 0

	for !win.Closed() {
		// Update the game
		switch gameState {
		case GameStateTitleScreen:
			if win.JustPressed(pixelgl.KeyEnter) {
				gameState = GameStatePlaying
			}
		case GameStatePlaying:
			playerLevelTxt.Clear()
			fmt.Fprintln(playerLevelTxt, "Player level:", player.Level)
			switch gameScreen {
			case GameScreenVillage:
				if win.JustPressed(pixelgl.KeyRight) {
					gameScreen = GameScreenForest
					isInQuest = true
					questStepsRemaining = 5
					fmt.Fprintln(questStepsRemainingTxt, "Quest steps remaining:", questStepsRemaining)
				}
			case GameScreenForest:
				if win.JustPressed(pixelgl.KeyRight) {
					questStepsRemaining -= 1
					if questStepsRemaining > 0 {
						questStepsRemainingTxt.Clear()
						fmt.Fprintln(questStepsRemainingTxt, "Quest steps remaining:", questStepsRemaining)
					} else {
						gameScreen = GameScreenVillage
					}
				}
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
				villagePromptTxt.Draw(win, pixel.IM)
			case GameScreenForest:
				backgroundForestSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
			}
			// Draw UI common to all screens
			if isInQuest == true {
				blackSprite.DrawColorMask(win, pixel.IM.ScaledXY(blackSprite.Picture().Bounds().Center(), pixel.V(18, 2)).Moved(pixel.V(1000, 738)), color.RGBA{255, 255, 255, 180})
				questStepsRemainingTxt.Draw(win, pixel.IM)
			}
			playerLevelTxt.Draw(win, pixel.IM)
		}

		win.Update()

		<-fps60
	}
}

func main() {
	pixelgl.Run(run)
}
