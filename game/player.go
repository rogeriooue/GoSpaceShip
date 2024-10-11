package game

import (
	"gospaceship/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfWidth := float64(bounds.Dx() / 2)

	position := Vector{
		X: (screenWidth / 2) - halfWidth,
		Y: 500,
	}

	return &Player{
		image:             image,
		game:              game,
		position:          position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}

	p.laserLoadingTimer.Update()

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()

		bounds := p.image.Bounds()
		halfWidth := float64(bounds.Dx() / 2)
		halfHeight := float64(bounds.Dy() / 2)

		spawnPos := Vector{
			p.position.X + halfWidth,
			p.position.Y - halfHeight/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLaser(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Position X and Y to draw the image
	op.GeoM.Translate(p.position.X, p.position.Y)

	// Draw the image on the screen
	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()

	return NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
