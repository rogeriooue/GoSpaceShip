package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gospaceship/assets"
	"math/rand"
)

type Meteor struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewMeteor() *Meteor {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	speed := rand.Float64() * 13

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	return &Meteor{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (m *Meteor) Update() {
	m.position.Y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Position X and Y to draw the image
	op.GeoM.Translate(m.position.X, m.position.Y)

	// Draw the image on the screen
	screen.DrawImage(m.image, op)
}

func (m *Meteor) Collider() Rect {
	bounds := m.image.Bounds()

	return NewRect(
		m.position.X,
		m.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
