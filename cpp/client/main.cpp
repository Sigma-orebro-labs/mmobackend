#include <SDL.h>
#include <memory>
#include <stdio.h>

#include "game/object.h"
#include "network/win_sock.h"
#include "network/tcp_client.h"
#include "network/udp_client.h"

bool init_sdl();

const int window_width = 800;
const int window_height = 800;

SDL_Window* window = nullptr;
SDL_Renderer* renderer = nullptr;

SDL_Rect enemy_rect = { 0, 0, 30, 30 };

int main(int argc, char* argv[])
{
	WinSock::initialize();

	auto tcp_client = std::make_unique<TcpClient>("localhost", "11000");
	tcp_client->create_connection();

	auto udp_client = std::make_unique<UdpClient>("127.0.0.1", 11001);
	udp_client->send();

	enemy_rect.x = 0;
	enemy_rect.y = 0;


	if (!init_sdl())
	{
		printf("SDL could not be initialized! SDL_Error: %s\n", SDL_GetError());
		return 1;
	}

	auto running = true;
	SDL_Event e;
	auto last_frame_time = SDL_GetTicks();

	auto player = std::make_unique<Object>(window_width / 2, window_height / 2, 30, 30);

	while (running)
	{
		while (SDL_PollEvent(&e) != 0)
		{
			if (e.type == SDL_QUIT)
			{
				running = false;
				continue;
			}

			player->handle_events(&e);
		}

		auto ticks = SDL_GetTicks();
		auto dt = ((ticks - last_frame_time) / 1000.f);
		last_frame_time = ticks;

		// Clear the window
		SDL_SetRenderDrawColor(renderer, 0xFF, 0xFF, 0xFF, 0xFF);
		SDL_RenderClear(renderer);

		player->update(dt);
		player->render(renderer);

		SDL_SetRenderDrawColor(renderer, 0xFF, 0, 0, 0xFF);
		SDL_RenderFillRect(renderer, &enemy_rect);

		SDL_RenderPresent(renderer);
	}

	SDL_DestroyRenderer(renderer);
	SDL_DestroyWindow(window);
	SDL_Quit();
	WinSock::cleanup();

	return 0;
}

bool init_sdl()
{
	if (SDL_Init(SDL_INIT_EVERYTHING) < 0)
	{
		return false;
	}
	else
	{
		window = SDL_CreateWindow("MMO", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, window_width, window_height, SDL_WINDOW_SHOWN);
		if (window == nullptr)
		{
			return false;
		}
		else
		{
			//Create a vsynced renderer for our window
			renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED | SDL_RENDERER_PRESENTVSYNC);
			if (renderer == nullptr)
			{
				return false;
			}
			else
			{
				SDL_SetRenderDrawColor(renderer, 0xFF, 0xFF, 0xFF, 0xFF);
			}
		}
	}

	return true;
}