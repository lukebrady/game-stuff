package main

import (
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	runtime.LockOSThread()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	// Create the SDL window.
	window, err := sdl.CreateWindow("Window Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		500, 319, sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	bmp, err := sdl.LoadBMP("./src/background.bmp")
	if err != nil {
		panic(err)
	}
	defer bmp.Free()
	go FadeInSurface(bmp, renderer, 0, time.Millisecond*100)
	go playMenuAudio()
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
