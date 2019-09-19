package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

const (
	width, height = 50, 10
	cellEmpty     = ' '
	cellBall      = '⚾'
	maxFrames     = 600
	speed         = time.Second / 20
)

func main() {
	var (
		cell   rune
		px, py int
		vx, vy = 1, 1
	)

	board := make([][]bool, width)

	for i := range board {
		board[i] = make([]bool, height)
	}

	buf := make([]rune, width*height)

	board[px][py] = true

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		board[px][py] = false
		px += vx
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		py += vy
		if py <= 0 || py >= height-1 {
			vy *= -1
			py += vy
		}
		board[px][py] = true

		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(speed)
	}
}
