package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	for pos, color := range theGrid {
		vector.DrawFilledRect(screen, halfGrid+float32(pos.X*int(tileSize)), halfGrid+float32(pos.Y*int(tileSize)), float32(tileSize-tileBorder), float32(tileSize-tileBorder), color, false)
	}
}
