package main

import (
	"image/color"
	"log"
	"math/rand"
	"sync"

	"github.com/hajimehoshi/ebiten"
)

const scale int = 2
const width = 300
const height = 300

var blue color.Color = color.RGBA{69, 145, 196, 255}
var yellow color.Color = color.RGBA{255, 230, 120, 255}
var grid [width][height]uint8 = [width][height]uint8{}
var buffer [width][height]uint8 = [width][height]uint8{}
var count int = 0

func update() error {
	var wg sync.WaitGroup
	works := 4
	rowsPerWorker := (height - 2) / works
	for i := 0; i < works; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			start := 1 + workerID*rowsPerWorker
			end := start + rowsPerWorker

			if workerID == works-1 {
				end = height - 1
			}
			for x := 1; x < width-1; x++ {
				for y := start; y < end; y++ {
					buffer[x][y] = 0
					n := grid[x-1][y-1] + grid[x-1][y+0] + grid[x-1][y+1] + grid[x+0][y-1] + grid[x+0][y+1] + grid[x+1][y-1] + grid[x+1][y+0] + grid[x+1][y+1]

					if grid[x][y] == 0 && n == 3 {
						buffer[x][y] = 1
					} else if n < 2 || n > 3 {
						buffer[x][y] = 0
					} else {
						buffer[x][y] = grid[x][y]
					}
				}
			}
		}(i)
	}
	wg.Wait()

	temp := buffer
	buffer = grid
	grid = temp
	return nil
}
func display(window *ebiten.Image) {
	window.Fill(blue)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for i := 0; i < scale; i++ {
				for j := 0; j < scale; j++ {
					if grid[x][y] == 1 {
						window.Set(x*scale+i, y*scale+j, yellow)
					}
				}
			}
		}
	}
}
func frame(window *ebiten.Image) error {
	count++
	var err error = nil
	if count == 20 {
		err = update()
		count = 0
	}
	if !ebiten.IsDrawingSkipped() {
		display(window)
	}

	return err
}
func main() {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if rand.Float32() < 0.5 {
				grid[x][y] = 1
			}
		}
	} //randomize grid

	if err := ebiten.Run(frame, width, height, 2, "Game of Life"); err != nil {
		log.Fatal(err)
	}
}
