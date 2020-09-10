package main

import (
	"math"

	"github.com/faiface/pixel/pixelgl"
)

func normalizeBatMove(batPosY float64, move float64) float64 {
	batHeight := 120.0
	batSpriteOffset := batHeight/2 + 22.0 // distance from sprite border

	halfScreenHeight := WINDOW_HEIGHT / 2
	distanceToUpperLowerBoundaries := halfScreenHeight - batSpriteOffset

	if math.Abs(batPosY-halfScreenHeight) > distanceToUpperLowerBoundaries {
		return 0.0 // can't move - against borders
	}

	xMin := batSpriteOffset
	xMax := WINDOW_HEIGHT - batSpriteOffset

	return math.Min(xMax, math.Max(xMin, batPosY+move)) - batPosY
}

func moveAi(ball ball, dt, batPosY float64) float64 {
	move := 0.0
	yTarget := 0.0
	isRightSide := ball.pos.X >= (WINDOW_WIDTH / 2)

	if !isRightSide {
		return move
	}

	if ball.mvmt.X == 1 { // ball moving right
		yTarget = ball.pos.Y // follows ball
	} else {
		yTarget = WINDOW_HEIGHT / 2 // return to middle
	}

	if batPosY > yTarget {
		move = -math.Min(AI_SPEED*dt, batPosY-yTarget)
	} else {
		move = math.Min(yTarget-batPosY, AI_SPEED*dt)
	}

	return normalizeBatMove(batPosY, move)
}

func moveBat(win *pixelgl.Window, dt float64, batPosY float64) float64 {
	move := 0.0
	if win.Pressed(pixelgl.KeyUp) {
		move = PLAYER_SPEED * dt
	}
	if win.Pressed(pixelgl.KeyDown) {
		move = -PLAYER_SPEED * dt
	}

	return normalizeBatMove(batPosY, move)
}
