package main

import (
	"math"
	"testing"
)

func TestMovePaddleMovesUpAndDown(t *testing.T) {
	paddle := Paddle{
		Position: Vector{X: 30, Y: 100},
		Width:    14,
		Height:   60,
		Speed:    5,
	}

	MovePaddle(&paddle, -1, 300)
	assertFloatEquals(t, paddle.Position.Y, 95)

	MovePaddle(&paddle, 1, 300)
	assertFloatEquals(t, paddle.Position.Y, 100)
}

func TestMovePaddleStaysInsideField(t *testing.T) {
	paddle := Paddle{
		Position: Vector{X: 30, Y: 31},
		Width:    14,
		Height:   60,
		Speed:    10,
	}

	MovePaddle(&paddle, -1, 300)
	assertFloatEquals(t, paddle.Position.Y, 30)

	paddle.Position.Y = 269
	MovePaddle(&paddle, 1, 300)
	assertFloatEquals(t, paddle.Position.Y, 270)
}

func TestMoveBallUsesVelocity(t *testing.T) {
	ball := Ball{
		Position: Vector{X: 100, Y: 200},
		Velocity: Vector{X: 5, Y: -2},
		Radius:   8,
	}

	MoveBall(&ball)

	assertFloatEquals(t, ball.Position.X, 105)
	assertFloatEquals(t, ball.Position.Y, 198)
}

func TestBounceOffTopWall(t *testing.T) {
	ball := Ball{
		Position: Vector{X: 100, Y: 5},
		Velocity: Vector{X: 4, Y: -3},
		Radius:   8,
	}

	collided := BounceOffHorizontalWalls(&ball, 300)

	if !collided {
		t.Fatal("expected collision with top wall")
	}
	assertFloatEquals(t, ball.Position.Y, 8)
	assertFloatEquals(t, ball.Velocity.Y, 3)
}

func TestBounceOffBottomWall(t *testing.T) {
	ball := Ball{
		Position: Vector{X: 100, Y: 296},
		Velocity: Vector{X: 4, Y: 3},
		Radius:   8,
	}

	collided := BounceOffHorizontalWalls(&ball, 300)

	if !collided {
		t.Fatal("expected collision with bottom wall")
	}
	assertFloatEquals(t, ball.Position.Y, 292)
	assertFloatEquals(t, ball.Velocity.Y, -3)
}

func TestBounceOffHorizontalWallsDoesNothingInsideField(t *testing.T) {
	ball := Ball{
		Position: Vector{X: 100, Y: 150},
		Velocity: Vector{X: 4, Y: 3},
		Radius:   8,
	}

	collided := BounceOffHorizontalWalls(&ball, 300)

	if collided {
		t.Fatal("did not expect wall collision")
	}
	assertFloatEquals(t, ball.Velocity.Y, 3)
}

func TestHasPaddleCollision(t *testing.T) {
	paddle := Paddle{
		Position: Vector{X: 50, Y: 100},
		Width:    10,
		Height:   60,
	}

	tests := []struct {
		name     string
		ball     Ball
		expected bool
	}{
		{
			name:     "ball overlaps paddle",
			ball:     Ball{Position: Vector{X: 58, Y: 100}, Radius: 5},
			expected: true,
		},
		{
			name:     "ball misses horizontally",
			ball:     Ball{Position: Vector{X: 80, Y: 100}, Radius: 5},
			expected: false,
		},
		{
			name:     "ball misses vertically",
			ball:     Ball{Position: Vector{X: 50, Y: 150}, Radius: 5},
			expected: false,
		},
		{
			name:     "touching edge counts as collision",
			ball:     Ball{Position: Vector{X: 60, Y: 100}, Radius: 5},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HasPaddleCollision(test.ball, paddle)
			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestBounceFromPaddleReversesHorizontalDirection(t *testing.T) {
	paddle := Paddle{
		Position: Vector{X: 50, Y: 100},
		Height:   80,
	}
	ball := Ball{
		Position: Vector{X: 60, Y: 100},
		Velocity: Vector{X: -5, Y: 1},
	}

	BounceFromPaddle(&ball, paddle)

	assertFloatEquals(t, ball.Velocity.X, 5)
	assertFloatEquals(t, ball.Velocity.Y, 1)
}

func TestBounceFromPaddleChangesVerticalDirectionBasedOnHitPosition(t *testing.T) {
	paddle := Paddle{
		Position: Vector{X: 50, Y: 100},
		Height:   80,
	}

	upperHit := Ball{
		Position: Vector{X: 60, Y: 80},
		Velocity: Vector{X: -5, Y: 0},
	}
	BounceFromPaddle(&upperHit, paddle)
	assertFloatEquals(t, upperHit.Velocity.Y, -1)

	lowerHit := Ball{
		Position: Vector{X: 60, Y: 120},
		Velocity: Vector{X: -5, Y: 0},
	}
	BounceFromPaddle(&lowerHit, paddle)
	assertFloatEquals(t, lowerHit.Velocity.Y, 1)
}

func TestDetectScore(t *testing.T) {
	tests := []struct {
		name     string
		ball     Ball
		expected ScoringPlayer
	}{
		{
			name:     "ball still inside",
			ball:     Ball{Position: Vector{X: 300, Y: 100}, Radius: 8},
			expected: NoScore,
		},
		{
			name:     "right player scores",
			ball:     Ball{Position: Vector{X: -9, Y: 100}, Radius: 8},
			expected: RightPlayer,
		},
		{
			name:     "left player scores",
			ball:     Ball{Position: Vector{X: 609, Y: 100}, Radius: 8},
			expected: LeftPlayer,
		},
		{
			name:     "ball partly outside left is not enough",
			ball:     Ball{Position: Vector{X: 4, Y: 100}, Radius: 8},
			expected: NoScore,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DetectScore(test.ball, 600)
			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func assertFloatEquals(t *testing.T, actual, expected float64) {
	t.Helper()
	if math.Abs(actual-expected) > 0.000001 {
		t.Fatalf("expected %.6f, got %.6f", expected, actual)
	}
}
