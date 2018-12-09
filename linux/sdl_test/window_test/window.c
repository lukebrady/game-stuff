#include <SDL.h>
#include <stdio.h>

int main() {
	// Declare a pointer and initialize SDL2.
	SDL_Window *window;
	SDL_Init(SDL_INIT_VIDEO);

	// Create the application window.
	window = SDL_CreateWindow(
		"My test SDL Window.",
		SDL_WINDOWPOS_UNDEFINED,
		SDL_WINDOWPOS_UNDEFINED,
		640,
		480,
		SDL_WINDOW_OPENGL
	);

	if(window == NULL) {
		printf("Could not create the SDL window.\n");
		return 1;
	}

	// Create the program loop here.
	
	SDL_Delay(3000);
	SDL_DestroyWindow(window);

	SDL_Quit();
	return 0;
}
