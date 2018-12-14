package main

import (
	"io/ioutil"
	"log"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

// Audio programming for Cafe Sudo.

func playAudio() {
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}
	defer mix.CloseAudio()

	// Load entire WAV data from file
	data, err := ioutil.ReadFile("/Users/ltbrady/Desktop/Ocean.wav")
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
