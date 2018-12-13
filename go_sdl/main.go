package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func playAudio() {
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}
	defer mix.CloseAudio()

	// Load entire WAV data from file
	data, err := ioutil.ReadFile("/Users/ltbrady/Desktop/Media/itat19.wav")
	if err != nil {
		log.Println(err)
	}

	// Load WAV from data (memory)
	chunk, err := mix.QuickLoadWAV(data)
	if err != nil {
		log.Println(err)
	}
	defer chunk.Free()

	// Play 4 times
	chunk.Play(1, 5)

	// Wait until it finishes playing
	for mix.Playing(-1) == 1 {
		sdl.Delay(16)
	}
}

func main() {
	runtime.LockOSThread()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	// Create the SDL window.
	window, err := sdl.CreateWindow("Window Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 500, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	bmp, err := sdl.LoadBMP("./linux/sdl_test/window_test/img/background.bmp")
	if err != nil {
		panic(err)
	}
	defer bmp.Free()

	texture, err := renderer.CreateTextureFromSurface(bmp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		panic(err)
	}
	defer texture.Destroy()

	src := sdl.Rect{0, 0, 512, 512}
	dst := sdl.Rect{100, 50, 512, 512}

	renderer.Clear()
	renderer.Copy(texture, &src, &dst)
	renderer.Present()

	sdl.Delay(2000)
	go playAudio()
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
