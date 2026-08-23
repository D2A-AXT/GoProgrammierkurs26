package main

import "testing"

func TestNextPhase(t *testing.T) {
	tests := []struct {
		input TrafficPhase
		want  TrafficPhase
	}{
		{EastGreen, EastYellow},
		{EastYellow, SouthGreen},
		{SouthGreen, SouthYellow},
		{SouthYellow, EastGreen},
	}

	for _, tt := range tests {
		if got := NextPhase(tt.input); got != tt.want {
			t.Fatalf("NextPhase(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestLightForDirection(t *testing.T) {
	tests := []struct {
		phase     TrafficPhase
		direction Direction
		want      LightState
	}{
		{EastGreen, East, Green},
		{EastGreen, South, Red},
		{EastYellow, East, Yellow},
		{EastYellow, South, Red},
		{SouthGreen, East, Red},
		{SouthGreen, South, Green},
		{SouthYellow, East, Red},
		{SouthYellow, South, Yellow},
	}

	for _, tt := range tests {
		if got := LightForDirection(tt.phase, tt.direction); got != tt.want {
			t.Fatalf(
				"LightForDirection(%v, %v) = %v, want %v",
				tt.phase,
				tt.direction,
				got,
				tt.want,
			)
		}
	}
}

func TestDistanceToStopLineEast(t *testing.T) {
	car := Car{X: 40, Direction: East}

	if got := DistanceToStopLine(car, 100); got != 60 {
		t.Fatalf("distance = %.1f, want 60", got)
	}

	car.X = 120

	if got := DistanceToStopLine(car, 100); got != -20 {
		t.Fatalf("distance = %.1f, want -20", got)
	}
}

func TestDistanceToStopLineSouth(t *testing.T) {
	car := Car{Y: 30, Direction: South}

	if got := DistanceToStopLine(car, 80); got != 50 {
		t.Fatalf("distance = %.1f, want 50", got)
	}
}

func TestHasSafeDistanceEast(t *testing.T) {
	car := Car{
		X:         100,
		Direction: East,
		Length:    20,
	}

	ahead := Car{
		X:         150,
		Direction: East,
		Length:    20,
	}

	if !HasSafeDistance(car, ahead, 20) {
		t.Error("50 units center distance should be safe for length 20 and gap 20")
	}

	ahead.X = 135

	if HasSafeDistance(car, ahead, 20) {
		t.Error("35 units center distance should be too close")
	}
}

func TestHasSafeDistanceSouth(t *testing.T) {
	car := Car{
		Y:         100,
		Direction: South,
		Length:    20,
	}

	ahead := Car{
		Y:         145,
		Direction: South,
		Length:    20,
	}

	if !HasSafeDistance(car, ahead, 15) {
		t.Error("expected safe distance")
	}

	ahead.Y = 125

	if HasSafeDistance(car, ahead, 15) {
		t.Error("expected distance to be unsafe")
	}
}

func TestHasSafeDistanceIgnoresOtherDirection(t *testing.T) {
	car := Car{X: 100, Y: 100, Direction: East, Length: 30}
	other := Car{X: 105, Y: 100, Direction: South, Length: 30}

	if !HasSafeDistance(car, other, 20) {
		t.Error("cars in different directions must not be treated as same-lane traffic")
	}
}

func TestCanCarMoveGreen(t *testing.T) {
	car := Car{ID: 1, X: 80, Direction: East, Length: 20}

	if !CanCarMove(car, []Car{car}, Green, 100, 10, 5) {
		t.Error("car should move on green")
	}
}

func TestCanCarMoveRedBeforeStopLine(t *testing.T) {
	car := Car{ID: 1, X: 80, Direction: East, Length: 20}

	if CanCarMove(car, []Car{car}, Red, 100, 10, 5) {
		t.Error("car must stop at red")
	}
}

func TestCanCarMoveAfterStopLineEvenOnRed(t *testing.T) {
	car := Car{ID: 1, X: 110, Direction: East, Length: 20}

	if !CanCarMove(car, []Car{car}, Red, 100, 10, 5) {
		t.Error("car already inside intersection must continue")
	}
}

func TestCanCarMoveYellowIfTooCloseToStop(t *testing.T) {
	car := Car{ID: 1, X: 97, Direction: East, Length: 20}

	if !CanCarMove(car, []Car{car}, Yellow, 100, 10, 5) {
		t.Error("car that reaches the line in this step should continue on yellow")
	}
}

func TestCanCarMoveYellowIfFarAway(t *testing.T) {
	car := Car{ID: 1, X: 80, Direction: East, Length: 20}

	if CanCarMove(car, []Car{car}, Yellow, 100, 10, 5) {
		t.Error("car far from stop line should stop on yellow")
	}
}

func TestCanCarMoveStopsForCarAhead(t *testing.T) {
	car := Car{
		ID:        1,
		X:         100,
		Direction: East,
		Length:    20,
	}
	ahead := Car{
		ID:        2,
		X:         130,
		Direction: East,
		Length:    20,
	}

	cars := []Car{car, ahead}

	if CanCarMove(car, cars, Green, 200, 15, 5) {
		t.Error("car must stop if another car is too close")
	}
}

func TestMoveCar(t *testing.T) {
	east := Car{X: 10, Y: 20, Direction: East}
	MoveCar(&east, 5)

	if east.X != 15 || east.Y != 20 {
		t.Fatalf("east car = (%.1f, %.1f), want (15,20)", east.X, east.Y)
	}

	south := Car{X: 10, Y: 20, Direction: South}
	MoveCar(&south, 7)

	if south.X != 10 || south.Y != 27 {
		t.Fatalf("south car = (%.1f, %.1f), want (10,27)", south.X, south.Y)
	}
}

func TestRemoveExitedCars(t *testing.T) {
	cars := []Car{
		{ID: 1, X: 50, Y: 50, Direction: East},
		{ID: 2, X: 120, Y: 20, Direction: East},
		{ID: 3, X: 20, Y: 130, Direction: South},
		{ID: 4, X: 20, Y: 80, Direction: South},
	}

	got := RemoveExitedCars(cars, 100, 100)

	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}

	if got[0].ID != 1 || got[1].ID != 4 {
		t.Fatalf("remaining cars = %+v", got)
	}
}

func TestCountWaitingCars(t *testing.T) {
	cars := []Car{
		{ID: 1, X: 80, Direction: East},
		{ID: 2, X: 120, Direction: East}, // already after line
		{ID: 3, Y: 70, Direction: South},
	}

	// SouthGreen -> East is red, South is green.
	got := CountWaitingCars(cars, SouthGreen, 100, 100)

	if got != 1 {
		t.Fatalf("waiting = %d, want 1", got)
	}
}
