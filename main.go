package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct{}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
}

func (g *game) Layout(w, h int) (int, int) {
	return ebiten.WindowSize()
}

func main() {
	g := &game{}
	ebiten.SetWindowSize(1280, 720) // has no effect on browser
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
