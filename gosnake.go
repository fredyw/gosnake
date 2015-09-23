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
	x1 int = 1
	y1 int = 0
	x2 int = 60
	y2 int = 20
)

func drawTopLine() {
	colorDefault := termbox.ColorDefault
	for i := x1; i <= x2; i++ {
		var c rune
		if i == x1 {
			c = '\u250c'
		} else if i == x2 {
			c = '\u2510'
		} else {
			c = '\u2500'
		}
		termbox.SetCell(i, y1, c, colorDefault, colorDefault)
	}
}

func drawLeftLine() {
	colorDefault := termbox.ColorDefault
	for i := y1 + 1; i <= y2; i++ {
		c := '\u2502'
		termbox.SetCell(x1, i, c, colorDefault, colorDefault)
	}
}

func drawRightLine() {
	colorDefault := termbox.ColorDefault
	for i := x1; i <= x2; i++ {
		var c rune
		if i == x1 {
			c = '\u2514'
		} else if i == x2 {
			c = '\u2518'
		} else {
			c = '\u2500'
		}
		termbox.SetCell(i, y2+1, c, colorDefault, colorDefault)
	}
}

func drawBottomLine() {
	colorDefault := termbox.ColorDefault
	for i := y1 + 1; i <= y2; i++ {
		c := '\u2502'
		termbox.SetCell(x2, i, c, colorDefault, colorDefault)
	}
}

func drawBox() {
	drawTopLine()
	drawLeftLine()
	drawRightLine()
	drawBottomLine()
}

func redrawAll() {
	colorDefault := termbox.ColorDefault
	termbox.Clear(colorDefault, colorDefault)

	drawBox()

	termbox.Flush()
}

func runGame() {
	err := termbox.Init()
	if err != nil {
		errorAndExit(err)
	}
	defer termbox.Close()
	redrawAll()
mainLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainLoop
			}
		}
		redrawAll()
	}
}

func errorAndExit(message interface{}) {
	fmt.Println(message)
	os.Exit(1)
}

func main() {
	runGame()
}
