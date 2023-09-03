package main

import (
	"image/color"
	"sync"
)

var (
	updateGameSizeLock sync.Mutex

	ScreenWidth, ScreenHeight int
	/* Game board values */
	boardPixels         = boardSize * gridSize
	gridSize    uint16  = uint16(ScreenHeight / boardSize)
	tileSize    uint16  = gridSize - tileBorder
	halfGrid    float32 = float32(gridSize) / 2.0

	theGrid  map[XY]color.Color
	gridLock sync.Mutex
)
