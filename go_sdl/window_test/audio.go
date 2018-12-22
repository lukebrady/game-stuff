package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

// Audio programming.
func oceanWaveAudio() {
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}
	defer mix.CloseAudio()

	// Load WAV from data (memory)
	chunk, err := mix.LoadWAV("./src/Ocean.wav")
	if err != nil {
		log.Println(err)
	}
	defer chunk.Free()

	chunk.Volume(50)

	// Play 4 times
	chunk.Play(1, 5)

	// Wait until it finishes playing
	for mix.Playing(-1) == 1 {
		sdl.Delay(16)
	}
}

// Audio programming.
func boatHornAudio() {
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}
	defer mix.CloseAudio()

	// Load WAV from data (memory)
	chunk, err := mix.LoadWAV("./src/Boat.wav")
	if err != nil {
		log.Println(err)
	}
	defer chunk.Free()

	chunk.Volume(25)

	// Play 4 times
	chunk.Play(1, 1)

	// Wait until it finishes playing
	for mix.Playing(-1) == 1 {
		sdl.Delay(16)
	}
}

func playMenuAudio() {
	go oceanWaveAudio()
	for x := 0; x < 100; x++ {
		time.Sleep(time.Millisecond * 100)
	}
	go boatHornAudio()
}
