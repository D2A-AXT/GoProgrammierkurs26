package main

import "testing"

func testWorld() World {
	return World{
		Cells: [][]rune{
			[]rune("#####"),
			[]rune("#   #"),
			[]rune("# # #"),
			[]rune("#   #"),
			[]rune("#####"),
		},
	}
}

func TestDirectionVector(t *testing.T) {
	tests := []struct {
		direction Direction
		expected  Position
	}{
		{Up, Position{X: 0, Y: -1}},
		{Right, Position{X: 1, Y: 0}},
		{Down, Position{X: 0, Y: 1}},
		{Left, Position{X: -1, Y: 0}},
	}

	for _, test := range tests {
		result := DirectionVector(test.direction)

		if result != test.expected {
			t.Errorf(
				"DirectionVector(%v) = %+v, expected %+v",
				test.direction,
				result,
				test.expected,
			)
		}
	}
}

func TestIsInside(t *testing.T) {
	world := testWorld()

	tests := []struct {
		position Position
		expected bool
	}{
		{Position{X: 0, Y: 0}, true},
		{Position{X: 4, Y: 4}, true},
		{Position{X: 2, Y: 2}, true},
		{Position{X: -1, Y: 0}, false},
		{Position{X: 0, Y: -1}, false},
		{Position{X: 5, Y: 0}, false},
		{Position{X: 0, Y: 5}, false},
	}

	for _, test := range tests {
		result := IsInside(world, test.position)

		if result != test.expected {
			t.Errorf(
				"IsInside(%+v) = %v, expected %v",
				test.position,
				result,
				test.expected,
			)
		}
	}
}

func TestIsWall(t *testing.T) {
	world := testWorld()

	if !IsWall(world, Position{X: 0, Y: 0}) {
		t.Error("expected (0,0) to be a wall")
	}

	if IsWall(world, Position{X: 1, Y: 1}) {
		t.Error("expected (1,1) to be free")
	}

	if !IsWall(world, Position{X: 2, Y: 2}) {
		t.Error("expected (2,2) to be a wall")
	}

	if !IsWall(world, Position{X: -1, Y: 1}) {
		t.Error("positions outside the world must count as walls")
	}
}

func TestCanMove(t *testing.T) {
	world := testWorld()
	robot := Robot{
		Position:  Position{X: 1, Y: 1},
		Direction: Right,
	}

	if CanMove(world, robot, Up) {
		t.Error("robot must not move into the upper wall")
	}

	if !CanMove(world, robot, Right) {
		t.Error("robot should be able to move right")
	}

	if !CanMove(world, robot, Down) {
		t.Error("robot should be able to move down")
	}
}

func TestMoveRobotMovesOnFreeCell(t *testing.T) {
	world := testWorld()
	robot := Robot{
		Position:  Position{X: 1, Y: 1},
		Direction: Up,
	}

	MoveRobot(world, &robot, Right)

	if robot.Position != (Position{X: 2, Y: 1}) {
		t.Fatalf("robot position = %+v, expected (2,1)", robot.Position)
	}

	if robot.Direction != Right {
		t.Errorf("robot direction = %v, expected Right", robot.Direction)
	}
}

func TestMoveRobotTurnsButDoesNotEnterWall(t *testing.T) {
	world := testWorld()
	robot := Robot{
		Position:  Position{X: 1, Y: 1},
		Direction: Right,
	}

	MoveRobot(world, &robot, Up)

	if robot.Position != (Position{X: 1, Y: 1}) {
		t.Fatalf("robot moved into a wall: %+v", robot.Position)
	}

	if robot.Direction != Up {
		t.Errorf("robot should still turn towards Up")
	}
}

func TestHasReachedGoal(t *testing.T) {
	goal := Position{X: 3, Y: 3}

	if !HasReachedGoal(
		Robot{Position: goal},
		goal,
	) {
		t.Error("robot standing on goal should have reached it")
	}

	if HasReachedGoal(
		Robot{Position: Position{X: 2, Y: 3}},
		goal,
	) {
		t.Error("robot next to goal must not count as finished")
	}
}

func TestTurnHelpers(t *testing.T) {
	if TurnRight(Up) != Right {
		t.Error("right of Up must be Right")
	}

	if TurnRight(Left) != Up {
		t.Error("right of Left must wrap to Up")
	}

	if TurnLeft(Up) != Left {
		t.Error("left of Up must wrap to Left")
	}

	if TurnBack(Right) != Left {
		t.Error("back of Right must be Left")
	}
}

func TestChooseDirectionUsesRightHandRule(t *testing.T) {
	world := World{
		Cells: [][]rune{
			[]rune("#####"),
			[]rune("#   #"),
			[]rune("##  #"),
			[]rune("#   #"),
			[]rune("#####"),
		},
	}

	// Robot schaut nach rechts.
	// Rechts davon (Down) ist frei, also muss er nach unten abbiegen.
	robot := Robot{
		Position:  Position{X: 2, Y: 1},
		Direction: Right,
	}

	result := ChooseDirection(world, robot)

	if result != Down {
		t.Errorf("expected Down, got %v", result)
	}
}

func TestChooseDirectionGoesStraightIfRightBlocked(t *testing.T) {
	world := World{
		Cells: [][]rune{
			[]rune("#####"),
			[]rune("#   #"),
			[]rune("#####"),
			[]rune("#   #"),
			[]rune("#####"),
		},
	}

	robot := Robot{
		Position:  Position{X: 2, Y: 1},
		Direction: Right,
	}

	result := ChooseDirection(world, robot)

	if result != Right {
		t.Errorf("expected Right, got %v", result)
	}
}
