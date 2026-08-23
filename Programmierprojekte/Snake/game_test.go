package main

import "testing"

func TestNewSnakeCreatesExpectedBody(t *testing.T) {
	snake := NewSnake(Point{X: 5, Y: 4}, 3, Right)

	expected := []Point{
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 3, Y: 4},
	}

	assertBodyEquals(t, snake.Body, expected)
}

func TestChangeDirectionAllowsTurn(t *testing.T) {
	snake := NewSnake(Point{X: 5, Y: 5}, 3, Right)

	ChangeDirection(&snake, Up)

	if snake.Direction != Up {
		t.Fatalf("expected direction %+v, got %+v", Up, snake.Direction)
	}
}

func TestChangeDirectionRejectsImmediateReverse(t *testing.T) {
	snake := NewSnake(Point{X: 5, Y: 5}, 3, Right)

	ChangeDirection(&snake, Left)

	if snake.Direction != Right {
		t.Fatalf("a snake moving right must not immediately turn left")
	}
}

func TestMoveSnakeMovesEveryBodyPart(t *testing.T) {
	snake := Snake{
		Body: []Point{
			{X: 5, Y: 4},
			{X: 4, Y: 4},
			{X: 3, Y: 4},
		},
		Direction: Right,
	}

	MoveSnake(&snake)

	expected := []Point{
		{X: 6, Y: 4},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
	}

	assertBodyEquals(t, snake.Body, expected)
}

func TestMoveSnakePreservesLength(t *testing.T) {
	snake := NewSnake(Point{X: 5, Y: 5}, 4, Down)

	MoveSnake(&snake)

	if len(snake.Body) != 4 {
		t.Fatalf("expected length 4 after moving, got %d", len(snake.Body))
	}
}

func TestHasWallCollision(t *testing.T) {
	tests := []struct {
		name     string
		head     Point
		width    int
		height   int
		expected bool
	}{
		{name: "inside", head: Point{X: 3, Y: 2}, width: 10, height: 8, expected: false},
		{name: "left wall", head: Point{X: -1, Y: 2}, width: 10, height: 8, expected: true},
		{name: "right wall", head: Point{X: 10, Y: 2}, width: 10, height: 8, expected: true},
		{name: "top wall", head: Point{X: 3, Y: -1}, width: 10, height: 8, expected: true},
		{name: "bottom wall", head: Point{X: 3, Y: 8}, width: 10, height: 8, expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			snake := Snake{Body: []Point{test.head}, Direction: Right}
			result := HasWallCollision(snake, test.width, test.height)
			if result != test.expected {
				t.Fatalf("expected %v, got %v for head %+v", test.expected, result, test.head)
			}
		})
	}
}

func TestHasSelfCollision(t *testing.T) {
	colliding := Snake{
		Body: []Point{
			{X: 4, Y: 3},
			{X: 4, Y: 2},
			{X: 3, Y: 2},
			{X: 3, Y: 3},
			{X: 4, Y: 3},
		},
	}

	if !HasSelfCollision(colliding) {
		t.Fatal("expected self collision")
	}

	notColliding := NewSnake(Point{X: 5, Y: 5}, 4, Right)
	if HasSelfCollision(notColliding) {
		t.Fatal("did not expect self collision")
	}
}

func TestHasEatenFood(t *testing.T) {
	snake := NewSnake(Point{X: 5, Y: 5}, 3, Right)

	if !HasEatenFood(snake, Point{X: 5, Y: 5}) {
		t.Fatal("expected food at the head to be eaten")
	}

	if HasEatenFood(snake, Point{X: 6, Y: 5}) {
		t.Fatal("did not expect food away from the head to be eaten")
	}
}

func TestGrowSnakeAddsOneBodyPart(t *testing.T) {
	snake := NewSnake(Point{X: 5, Y: 5}, 3, Right)
	oldTail := snake.Body[len(snake.Body)-1]

	GrowSnake(&snake)

	if len(snake.Body) != 4 {
		t.Fatalf("expected length 4, got %d", len(snake.Body))
	}

	if snake.Body[len(snake.Body)-1] != oldTail {
		t.Fatalf("expected new segment at old tail position %+v, got %+v", oldTail, snake.Body[len(snake.Body)-1])
	}
}

func assertBodyEquals(t *testing.T, actual, expected []Point) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Fatalf("expected body length %d, got %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Fatalf("segment %d: expected %+v, got %+v", i, expected[i], actual[i])
		}
	}
}
