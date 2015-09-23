// The MIT License (MIT)
//
// Copyright (c) 2015 Fredy Wijaya
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

const (
	topLeftX     int           = 1
	topLeftY     int           = 0
	bottomRightX int           = 60
	bottomRightY int           = 20
	snakeX       int           = bottomRightX / 2
	snakeY       int           = bottomRightY / 2
	idle         int           = -1
	left         int           = 0
	right        int           = 1
	up           int           = 2
	down         int           = 3
	speed        time.Duration = 500
	xStep        int           = 2
	yStep        int           = 1
)

func drawTopLine() {
	colorDefault := termbox.ColorDefault
	for i := topLeftX; i <= bottomRightX; i++ {
		var c rune
		if i == topLeftX {
			c = '\u250c'
		} else if i == bottomRightX {
			c = '\u2510'
		} else {
			c = '\u2500'
		}
		termbox.SetCell(i, topLeftY, c, colorDefault, colorDefault)
	}
}

func drawLeftLine() {
	colorDefault := termbox.ColorDefault
	for i := topLeftY + 1; i <= bottomRightY; i++ {
		c := '\u2502'
		termbox.SetCell(topLeftX, i, c, colorDefault, colorDefault)
	}
}

func drawRightLine() {
	colorDefault := termbox.ColorDefault
	for i := topLeftX; i <= bottomRightX; i++ {
		var c rune
		if i == topLeftX {
			c = '\u2514'
		} else if i == bottomRightX {
			c = '\u2518'
		} else {
			c = '\u2500'
		}
		termbox.SetCell(i, bottomRightY+1, c, colorDefault, colorDefault)
	}
}

func drawBottomLine() {
	colorDefault := termbox.ColorDefault
	for i := topLeftY + 1; i <= bottomRightY; i++ {
		c := '\u2502'
		termbox.SetCell(bottomRightX, i, c, colorDefault, colorDefault)
	}
}

func drawBox() {
	drawTopLine()
	drawLeftLine()
	drawRightLine()
	drawBottomLine()
}

func drawSnake(snake *snake) {
	colorDefault := termbox.ColorDefault
	for _, coordinate := range snake.coordinates {
		termbox.SetCell(coordinate.x, coordinate.y, '@', colorDefault, colorDefault)
	}
}

func drawText(text string) {
	colorDefault := termbox.ColorDefault
	x := topLeftX + 1
	y := bottomRightY + 2
	for _, ch := range text {
		termbox.SetCell(x, y, ch, colorDefault, colorDefault)
		x++
	}
}

func drawCoordinates(snake *snake) {
	text := fmt.Sprintf("x=%d, y=%d", snake.coordinates[0].x, snake.coordinates[0].y)
	drawText(text)
}

func drawFood(food *food) {
	colorDefault := termbox.ColorDefault
	for _, coordinate := range food.coordinates {
		termbox.SetCell(coordinate.x, coordinate.y, '\u2665', colorDefault, colorDefault)
	}
}

func redrawAll(snake *snake, food *food) {
	colorDefault := termbox.ColorDefault
	termbox.Clear(colorDefault, colorDefault)

	drawBox()
	drawSnake(snake)
	drawFood(food)
	drawCoordinates(snake)

	termbox.Flush()
}

type food struct {
	coordinates []coordinate
}

type coordinate struct {
	x int
	y int
}

type snake struct {
	coordinates []coordinate
	direction   int
}

func (s *snake) update(moveHead func(idx int)) {
	// move the head
	idx := 0
	prev := s.coordinates[idx]
	moveHead(idx)
	// update the tail
	for i := idx + 1; i < len(s.coordinates); i++ {
		tmp := s.coordinates[i]
		s.coordinates[i].x = prev.x
		s.coordinates[i].y = prev.y
		prev = tmp
	}
}

func (s *snake) moveUpIdx(idx int) {
	s.coordinates[idx].y -= yStep
	if s.coordinates[idx].y <= topLeftY {
		s.coordinates[idx].y = bottomRightY
	}
	s.direction = up
}

func (s *snake) moveUp() {
	s.update(s.moveUpIdx)
}

func (s *snake) moveDownIdx(idx int) {
	s.coordinates[idx].y += yStep
	if s.coordinates[idx].y >= bottomRightY {
		s.coordinates[idx].y = topLeftY + yStep
	}
	s.direction = down
}

func (s *snake) moveDown() {
	s.update(s.moveDownIdx)
}

func (s *snake) moveLeftIdx(idx int) {
	s.coordinates[idx].x -= xStep
	if s.coordinates[idx].x <= topLeftX+1 {
		s.coordinates[idx].x = bottomRightX - xStep
	}
	s.direction = left
}

func (s *snake) moveLeft() {
	s.update(s.moveLeftIdx)
}

func (s *snake) moveRightIdx(idx int) {
	s.coordinates[idx].x += xStep
	if s.coordinates[idx].x >= bottomRightX-1 {
		s.coordinates[idx].x = topLeftX + xStep
	}
	s.direction = right
}

func (s *snake) moveRight() {
	s.update(s.moveRightIdx)
}

func (s *snake) move(direction int) {
	if s.direction == idle {
		if direction == left {
			s.moveLeft()
		} else if direction == right {
			s.moveRight()
		} else if direction == up {
			s.moveUp()
		} else if direction == down {
			s.moveDown()
		}
	} else {
		// the snake can't go backward
		if direction == left && !(s.direction == left || s.direction == right) {
			s.moveLeft()
		} else if direction == right && !(s.direction == left || s.direction == right) {
			s.moveRight()
		} else if direction == up && !(s.direction == up || s.direction == down) {
			s.moveUp()
		} else if direction == down && !(s.direction == up || s.direction == down) {
			s.moveDown()
		}
	}
}

func (s *snake) autoMove(food *food) {
	if s.direction == left {
		s.moveLeft()
	} else if s.direction == right {
		s.moveRight()
	} else if s.direction == up {
		s.moveUp()
	} else if s.direction == down {
		s.moveDown()
	}
	// TODO: refactor this
	head := s.coordinates[0]
	var newFood []coordinate
	for _, foodCoord := range food.coordinates {
		if head.x == foodCoord.x && head.y == foodCoord.y {
			s.coordinates = append([]coordinate{{x: head.x, y: head.y}}, s.coordinates...)
		} else {
			newFood = append(newFood, foodCoord)
		}
	}
	food.coordinates = newFood
}

func runGame() {
	err := termbox.Init()
	if err != nil {
		errorAndExit(err)
	}
	defer termbox.Close()
	// TODO: fix initial snake position
	snake := &snake{
		coordinates: []coordinate{
			coordinate{x: snakeX, y: snakeY},
			coordinate{x: snakeX - xStep, y: snakeY},
			coordinate{x: snakeX - (xStep * 2), y: snakeY},
		},
		direction: idle,
	}
	// TODO: randomize the food
	food := &food{
		coordinates: []coordinate{
			coordinate{x: 20, y: 10},
			coordinate{x: 30, y: 14},
		},
	}
	redrawAll(snake, food)

	ticker := time.NewTicker(speed * time.Millisecond)

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()
loop:
	for {
		select {
		// TODO: bug for speed
		case ev := <-eventQueue:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowUp:
				snake.move(up)
			case termbox.KeyArrowDown:
				snake.move(down)
			case termbox.KeyArrowLeft:
				snake.move(left)
			case termbox.KeyArrowRight:
				snake.move(right)
			}
		case <-ticker.C:
			snake.autoMove(food)
		}
		redrawAll(snake, food)
	}
}

func errorAndExit(message interface{}) {
	fmt.Println(message)
	os.Exit(1)
}

func main() {
	runGame()
}
