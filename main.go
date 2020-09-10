package main

import (
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	WINDOW_HEIGHT = 480.0
	WINDOW_WIDTH  = 800.0
	BALL_SPEED    = 250.0
	PLAYER_SPEED  = 500.0
	AI_SPEED      = 200.0
)

var (
	hitCount = 0
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "PonGo!",
		Bounds: pixel.R(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	bgSprite := loadSprite("images/table.png")
	ballSprite := loadSprite("images/ball.png")
	leftBatSprite := loadSprite("images/bat00.png")
	rightBatSprite := loadSprite("images/bat10.png")

	center := win.Bounds().Center()

	ball := ball{center, pixel.V(-1, 0)}
	leftBatPos := pixel.V(40, center.Y)
	rightBatPos := pixel.V(760, center.Y)

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		// slow motion with tab
		if win.Pressed(pixelgl.KeyTab) {
			dt /= 8
		}

		bgSprite.Draw(win, pixel.IM.Moved(center))
		leftBatSprite.Draw(win, pixel.IM.Moved(leftBatPos))
		rightBatSprite.Draw(win, pixel.IM.Moved(rightBatPos))
		ballSprite.Draw(win, pixel.IM.Moved(ball.pos))

		ball.move(center, leftBatPos, rightBatPos, dt)

		leftBatPos.Y += moveBat(win, dt, leftBatPos.Y)
		rightBatPos.Y += moveAi(ball, dt, rightBatPos.Y)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
