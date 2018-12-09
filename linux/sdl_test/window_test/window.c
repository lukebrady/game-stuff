#include <SDL.h>
#include <stdio.h>

int main() {
	// Declare a pointer and initialize SDL2.
	SDL_Window *window = NULL;
	SDL_Renderer *renderer = NULL;
	SDL_Texture *bitmapTex = NULL;
	SDL_Surface *bitmapSurface = NULL;
	SDL_Init(SDL_INIT_VIDEO);

	// Create the application window.
	window = SDL_CreateWindow(
		"My test SDL Window.",
		SDL_WINDOWPOS_UNDEFINED,
		SDL_WINDOWPOS_UNDEFINED,
		500,
		319,
		SDL_WINDOW_OPENGL
	);

	if(window == NULL) {
		printf("Could not create the SDL window.\n");
		return 1;
	}

	// Create the renderer.
	renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED);
	if(renderer == NULL) {
		printf("Could not create the SDL renderer.\n");
		return 1;
	}
	// Load the background image.
	bitmapSurface = SDL_LoadBMP("img/background.bmp");
	bitmapTex = SDL_CreateTextureFromSurface(renderer, bitmapSurface);
	SDL_FreeSurface(bitmapSurface);

	// Create the program loop here.
	while(1) {
		SDL_Event e;
		if(SDL_PollEvent(&e)) {
			if(e.type == SDL_QUIT) {
				break;
			}
		}

		SDL_RenderClear(renderer);
		SDL_RenderCopy(renderer, bitmapTex, NULL, NULL);
		SDL_RenderPresent(renderer);
	}

	// Destroy the texture, renderer, and window.
	SDL_DestroyTexture(bitmapTex);
	SDL_DestroyRenderer(renderer);
	SDL_DestroyWindow(window);

	SDL_Quit();

	return 0;
}
