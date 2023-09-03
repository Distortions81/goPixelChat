package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/irc"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	defaultWindowWidth  = 1280
	defaultWindowHeight = 720

	/* Game tiles */
	tileBorder = 1  //1px spacing
	boardSize  = 64 //Default board size

)

var (
	ScreenWidth, ScreenHeight int
	/* Game board values */
	boardPixels         = boardSize * gridSize
	gridSize    uint16  = uint16(ScreenHeight / boardSize)
	tileSize    uint16  = gridSize - tileBorder
	halfGrid    float32 = float32(gridSize) / 2.0
)

func main() {
	writer := &irc.Conn{}

	//Get aiuth
	auth, err := os.ReadFile("auth.txt")
	if err != nil {
		log.Fatal(err)
	}

	if 1 == 2 {
		//Connect
		writer.SetLogin("xboxtv81", "oauth:"+string(auth))
		if err := writer.Connect(); err != nil {
			panic("failed to start writer")
		}

		reader := twitch.IRC()
		reader.OnShardReconnect(onShardReconnect)
		reader.OnShardLatencyUpdate(onShardLatencyUpdate)
		reader.OnShardMessage(onShardMessage)

		if err := reader.Join("xboxtv81"); err != nil {
			panic(err)
		}
		fmt.Println("Connected to IRC!")
	}

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

func onShardReconnect(shardID int) {
	fmt.Printf("Shard #%d reconnected\n", shardID)
}

func onShardLatencyUpdate(shardID int, latency time.Duration) {
	fmt.Printf("Shard #%d has %dms ping\n", shardID, latency.Milliseconds())
}

func onShardMessage(shardID int, msg irc.ChatMessage) {
	fmt.Printf("#%s %s: %s\n", msg.Channel, msg.Sender.DisplayName, msg.Text)
}

func (g *Game) Update() error {
	return nil
}

var theGrid map[XY]color.Color

func (g *Game) Draw(screen *ebiten.Image) {
	for pos, color := range theGrid {
		vector.DrawFilledRect(screen, halfGrid+float32(pos.X*int(tileSize)), halfGrid+float32(pos.Y*int(tileSize)), float32(tileSize-tileBorder), float32(tileSize-tileBorder), color, false)
	}
}

func newGame() *Game {
	theGrid = make(map[XY]color.Color)
	newColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	theGrid[XY{X: 1, Y: 1}] = newColor
	theGrid[XY{X: 1, Y: 2}] = newColor
	theGrid[XY{X: 1, Y: 3}] = newColor

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

var updateGameSizeLock sync.Mutex

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

type Game struct {
}

type XY struct {
	X int
	Y int
}
