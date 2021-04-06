package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenSize int32 = 500
	cellSize   int32 = 10
	cellNumber int32 = screenSize / cellSize
)

func main() {
	// ---------- parse flags
	f := flag.String("pattern", ".", "name of the pattern")
	flag.Parse()
	if *f == "." {
		fmt.Println("parse flags:", "pattern missed")
		return
	}

	// ---------- open pattern file
	reader, err := os.OpenFile("./patterns/"+*f, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("open pattern file:", err)
		return
	}

	// ---------- read pattern file
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("read pattern file:", err)
		return
	}

	// ---------- initializing SDL
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	// ---------- creating SDL
	window, err := sdl.CreateWindow(
		"Game of life",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenSize,
		screenSize,
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		fmt.Println("creating window:", err)
		return
	}
	defer window.Destroy()

	// ---------- creating renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("creating renderer:", err)
		return
	}
	defer renderer.Destroy()

	// ---------- generating cells
	cells, err := cells_generator(data)
	if err != nil {
		fmt.Println("generating cells:", err)
		return
	}

	// ---------- game loop
	for {
		// ---------- user inputs
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		// ---------- cleaning screen
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		// ---------- print living cells
		renderer.SetDrawColor(0, 0, 0, 0)
		for y_index, y := range cells {
			y_position := int32(y_index) * cellSize
			for x_index, x := range y {
				x_position := int32(x_index) * cellSize
				if x {
					renderer.FillRect(&sdl.Rect{X: int32(x_position), Y: y_position, W: cellSize, H: cellSize})
				}
			}
		}

		// ---------- show the new screen
		renderer.Present()

		// ---------- calculating the next pattern after a pause
		time.Sleep(100 * time.Millisecond)
		next(&cells)
	}
}
