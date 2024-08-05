package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gospaceship/assets"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bounds := image.Bounds()
	halfWidth := float64(bounds.Dx() / 2)
	halfHeight := float64(bounds.Dy() / 2)

	position.X -= halfWidth
	position.Y -= halfHeight

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 7.0
	l.position.Y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Position X and Y to draw the image
	op.GeoM.Translate(l.position.X, l.position.Y)

	// Draw the image on the screen
	screen.DrawImage(l.image, op)
}

func (l *Laser) Collider() Rect {
	bounds := l.image.Bounds()

	return NewRect(
		l.position.X,
		l.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
