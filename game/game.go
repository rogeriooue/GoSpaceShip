package game

import (
	"fmt"
	"gospaceship/assets"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	player            *Player
	lasers            []*Laser
	meteors           []*Meteor
	stars             []*Star
	meteorsSpawnTimer *Timer
	starsSpawnTimer   *Timer
	score             int
	gameOver          bool
	canReset          bool
	gameOverTimer     time.Time
}

func NewGame() *Game {
	g := &Game{
		meteorsSpawnTimer: NewTimer(24),
		starsSpawnTimer:   NewTimer(48),
	}
	player := NewPlayer(g)
	g.player = player

	return g
}

// Update 60 frames per second
// 1 Tick = 1 x seconds
// Responsible for initializing the game
func (g *Game) Update() error {
	if g.gameOver {
		if time.Since(g.gameOverTimer) > 1*time.Second {
			g.canReset = true
		}
		if g.canReset && ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Reset()
		}
		return nil
	}

	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

	g.meteorsSpawnTimer.Update()

	if g.meteorsSpawnTimer.IsReady() {
		g.meteorsSpawnTimer.Reset()
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			fmt.Println("Game Over")
			g.gameOver = true
			g.canReset = false
			g.gameOverTimer = time.Now()
		}
	}

	for i, m := range g.meteors {
		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score += 10
			}
		}
	}

	g.starsSpawnTimer.Update()

	if g.starsSpawnTimer.IsReady() {
		g.starsSpawnTimer.Reset()
		s := NewStar()
		g.stars = append(g.stars, s)
	}

	for _, s := range g.stars {
		s.Update()
	}

	return nil
}

// Draw 60 frames per second
// Responsible for drawing the game on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
		text.Draw(screen, fmt.Sprintf("Score: %d", g.score), assets.FontUi, 20, 50, color.White)
		text.Draw(screen, fmt.Sprintf("Game Over"), assets.FontUi, 270, 300, color.White)
		text.Draw(screen, fmt.Sprintf("Press ENTER to Restart"), assets.FontUi, 80, 400, color.White)
		return
	}

	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, s := range g.stars {
		s.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), assets.FontUi, 20, 50, color.White)
}

// Layout Responsible for setting the screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLaser(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.lasers = nil
	g.meteors = nil
	g.stars = nil
	g.meteorsSpawnTimer.Reset()
	g.score = 0
	g.gameOver = false
}
