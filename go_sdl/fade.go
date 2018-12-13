package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// FadeInSurface fades in the supplied surface.
func FadeInSurface(bmp *sdl.Surface, renderer *sdl.Renderer, alpha float32, fadeSpeed float32, timeSpan float32) error {
	// Set the alpha starting point.
	for alpha < 255 {
		bmp.SetAlphaMod(uint8(alpha))
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
		alpha += fadeSpeed * timeSpan
		println(alpha)
	}
	return nil
}
