package main

import (
	"math"

	"github.com/faiface/pixel"
)

type ball struct {
	pos  pixel.Vec
	mvmt pixel.Vec
}

func (b *ball) move(center, leftBatPos, rightBatPos pixel.Vec, dt float64) {
	if b.mvmt.X > 0 { // ball going to the right
		b.processBatCollision(rightBatPos)
	} else {
		b.processBatCollision(leftBatPos)
	}

	if b.pos.X < leftBatPos.X || b.pos.X > rightBatPos.X { // reset when loosing
		b.pos.X = center.X
		b.pos.Y = center.Y
		b.mvmt.X = 1
		b.mvmt.Y = 0.0
		hitCount = 0
	}

	// bounce upper and lower borders
	border := 20.0
	if math.Abs(b.pos.Y-(WINDOW_HEIGHT/2)) > (WINDOW_HEIGHT/2)-border {
		b.mvmt.Y = -b.mvmt.Y
	}

	currentSpeed := BALL_SPEED + float64(hitCount)*20
	b.pos.X += b.mvmt.X * currentSpeed * dt
	b.pos.Y += b.mvmt.Y * currentSpeed * dt
}

func (b *ball) processBatCollision(batPos pixel.Vec) {
	xCol := float64(40 + 9) // distance from screen edge + half of width of the bat
	if math.Abs(b.pos.X-(WINDOW_WIDTH/2)) <= (WINDOW_WIDTH/2)-xCol {
		return // no X collision
	}

	yDelta := b.pos.Y - batPos.Y
	batHeight := 128.0
	if math.Abs(yDelta) < batHeight/2 { // Y check
		b.mvmt = pixel.V(
			-b.mvmt.X,
			yDelta/128, // {-0,5 ; 0.5 }
		)
		hitCount++
	}
}
