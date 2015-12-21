#include <SDL.h>
#include <memory>
#include <stdio.h>

#include "game/object.h"
#include "network/win_sock.h"
#include "network/udp_client.h"
#include "network/message.h"

bool init_sdl();
void update_player_position(const std::unique_ptr<UdpClient>& c, const std::unique_ptr<Object>& p);

const int window_width = 800;
const int window_height = 800;

SDL_Window* window = nullptr;
SDL_Renderer* renderer = nullptr;

SDL_Rect enemy_rect = { 0, 0, 30, 30 };

int main(int argc, char* argv[])
{
	WinSock::initialize();

	auto udp_client = std::make_unique<UdpClient>("127.0.0.1", 11001);

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
	update_player_position(udp_client, player);
	auto last_player_position_update = SDL_GetTicks();

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

		auto now = SDL_GetTicks();
		// Update our position every 50 ms
		if (now - last_player_position_update > 50)
		{
			update_player_position(udp_client, player);
			last_player_position_update = now;
		}
	}

	SDL_DestroyRenderer(renderer);
	SDL_DestroyWindow(window);
	SDL_Quit();
	WinSock::cleanup();

	return 0;
}

void update_player_position(const std::unique_ptr<UdpClient>& c, const std::unique_ptr<Object>& p)
{
	auto msg = Message(Command::UPDATE_PLAYER_POSITION_COMMAND);
	msg.append_to_body(static_cast<uint16_t>(p->get_x()));
	msg.append_to_body(static_cast<uint16_t>(p->get_y()));

	c->send(msg);
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