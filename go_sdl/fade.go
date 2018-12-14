package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// FadeInSurface fades in the supplied surface.
func FadeInSurface(bmp *sdl.Surface, renderer *sdl.Renderer, alpha uint8, timeSpan int) error {
	// Set the alpha starting point.
	for x := 0; x < 50; x++ {
		bmp.SetAlphaMod(alpha)
		texture, err := renderer.CreateTextureFromSurface(bmp)
		if err != nil {
			return err
		}
		defer texture.Destroy()

		src := sdl.Rect{0, 0, 500, 319}
		dst := sdl.Rect{0, 0, 500, 319}
		renderer.Clear()
		renderer.Copy(texture, &src, &dst)
		renderer.Present()
		alpha += 5
		time.Sleep(time.Millisecond * 100)
		println(alpha)
		if alpha > 255 {
			break
		}
	}
	return nil
}
