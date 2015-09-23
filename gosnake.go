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
)

const (
	topLeftX     int = 1
	topLeftY     int = 0
	bottomRightX int = 60
	bottomRightY int = 20
	snakeX           = bottomRightX / 2
	snakeY           = bottomRightY / 2
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
	x int
	y int
}

func (s *snake) moveUp() {
	s.y--
	if s.y <= topLeftY {
		s.y = bottomRightY
	}
}

func (s *snake) moveDown() {
	s.y++
	if s.y >= bottomRightY {
		s.y = topLeftY + 1
	}
}

func (s *snake) moveLeft() {
	s.x--
	if s.x <= topLeftX+1 {
		s.x = bottomRightX - 2
	}
}

func (s *snake) moveRight() {
	s.x++
	if s.x >= bottomRightX-1 {
		s.x = topLeftX + 2
	}
}

func runGame() {
	err := termbox.Init()
	if err != nil {
		errorAndExit(err)
	}
	defer termbox.Close()
	snake := snake{x: snakeX, y: snakeY}
	redrawAll(snake)
mainLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainLoop
			case termbox.KeyArrowUp:
				snake.moveUp()
			case termbox.KeyArrowDown:
				snake.moveDown()
			case termbox.KeyArrowLeft:
				snake.moveLeft()
			case termbox.KeyArrowRight:
				snake.moveRight()
			}
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
