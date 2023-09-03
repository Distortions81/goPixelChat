package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func startEbiten() {
	// Set up ebiten
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	/* We manaually clear, so we aren't forced to draw every frame */
	ScreenWidth, ScreenHeight = defaultWindowWidth, defaultWindowHeight

	/* Set up our window */
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("goPixelChat")

	/* Start ebiten */
	if err := ebiten.RunGameWithOptions(newGame(), nil); err != nil {
		return
	}
}

func updateGameSize() {

	/* Resize everything for the new window size */
	updateGameSizeLock.Lock()
	defer updateGameSizeLock.Unlock()

	gridSize = uint16(ScreenHeight / (boardSize + 1))
	if gridSize < 3 {
		gridSize = 3
	}
	tileSize = gridSize - tileBorder

	boardPixels = boardSize * gridSize
	if boardPixels < (boardSize+1)*3 {
		boardPixels = (boardSize + 1) * 3
	}
	halfGrid = float32(gridSize) / 2.0
}

func newGame() *Game {
	theGrid = make(map[XY]color.Color)

	theGrid[XY{X: 1, Y: 1}] = ColorRed
	theGrid[XY{X: 2, Y: 1}] = ColorGreen
	theGrid[XY{X: 3, Y: 1}] = ColorBlue

	updateGameSize()
	return &Game{}
}

/* Window size chaged, handle it */
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	if outsideWidth != ScreenWidth || outsideHeight != ScreenHeight {

		ScreenWidth, ScreenHeight = outsideWidth, outsideHeight
		updateGameSize()
	}

	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	return nil
}
