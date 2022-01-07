package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	DOWN         = iota
	LEFT         = iota * 64
	RIGHT        = iota * 64
	UP           = iota * 64
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 0
	frameWidth  = 64
	frameHeight = 64
	frameNum    = 4
)

var (
	state       = 0
	runnerImage *ebiten.Image
)

var (
	x = float64(screenWidth / 2)
	y = float64(screenHeight / 2)
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		y += 10
		state = DOWN
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		y -= 10
		state = UP
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		x += 10
		state = RIGHT
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		x -= 10
		state = LEFT
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(x, y)
	i := (g.count / 20) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY+state
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	f, err := os.ReadFile("ethanwalk.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
