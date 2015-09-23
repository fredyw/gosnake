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
	speed        time.Duration = 200
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

func drawSnake(snake snake) {
	colorDefault := termbox.ColorDefault
	termbox.SetCell(snake.x, snake.y, '@', colorDefault, colorDefault)
}

func redrawAll(snake snake) {
	colorDefault := termbox.ColorDefault
	termbox.Clear(colorDefault, colorDefault)

	drawBox()
	drawSnake(snake)

	termbox.Flush()
}

type snake struct {
	x         int
	y         int
	direction int
}

func (s *snake) moveUp() {
	s.y--
	if s.y <= topLeftY {
		s.y = bottomRightY
	}
	s.direction = up
}

func (s *snake) moveDown() {
	s.y++
	if s.y >= bottomRightY {
		s.y = topLeftY + 1
	}
	s.direction = down
}

func (s *snake) moveLeft() {
	s.x -= 2
	if s.x <= topLeftX+1 {
		s.x = bottomRightX - 2
	}
	s.direction = left
}

func (s *snake) moveRight() {
	s.x += 2
	if s.x >= bottomRightX-1 {
		s.x = topLeftX + 2
	}
	s.direction = right
}

func (s *snake) move() {
	if s.direction == left {
		s.moveLeft()
	} else if s.direction == right {
		s.moveRight()
	} else if s.direction == up {
		s.moveUp()
	} else if s.direction == down {
		s.moveDown()
	}
}

func runGame() {
	err := termbox.Init()
	if err != nil {
		errorAndExit(err)
	}
	defer termbox.Close()
	snake := snake{x: snakeX, y: snakeY, direction: idle}
	redrawAll(snake)

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
		case ev := <-eventQueue:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowUp:
				if snake.direction != up {
					snake.moveUp()
				}
			case termbox.KeyArrowDown:
				if snake.direction != down {
					snake.moveDown()
				}
			case termbox.KeyArrowLeft:
				if snake.direction != left {
					snake.moveLeft()
				}
			case termbox.KeyArrowRight:
				if snake.direction != right {
					snake.moveRight()
				}
			}
		case <-ticker.C:
			snake.move()
		}
		redrawAll(snake)
	}
}

func errorAndExit(message interface{}) {
	fmt.Println(message)
	os.Exit(1)
}

func main() {
	runGame()
}
