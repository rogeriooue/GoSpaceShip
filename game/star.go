package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gospaceship/assets"
	"math/rand"
)

type Star struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewStar() *Star {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]

	speed := rand.Float64() * 6

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	return &Star{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (s *Star) Update() {
	s.position.Y += s.speed
}

func (s *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Position X and Y to draw the image
	op.GeoM.Translate(s.position.X, s.position.Y)

	// Draw the image on the screen
	screen.DrawImage(s.image, op)
}
