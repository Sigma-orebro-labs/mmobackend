#include "object.h"

Object::Object(int x, int y, int width, int height)
	:
	vel_x_(0),
	vel_y_(0)
{
	body_ = { x, y, width, height };
}

void Object::render(SDL_Renderer* renderer) const
{
	SDL_SetRenderDrawColor(renderer, 0, 0, 0, 0xFF);
	SDL_RenderFillRect(renderer, &body_);
}

void Object::update(float dt)
{
	body_.x += static_cast<int>(dt * vel_x_);
	body_.y += static_cast<int>(dt * vel_y_);
}

void Object::handle_events(const SDL_Event* e)
{
	auto speed = 100;

	//A key was pushed
	if (e->type == SDL_KEYDOWN && e->key.repeat == 0)
	{
		switch (e->key.keysym.sym)
		{
			case SDLK_UP: vel_y_ = -100; break;
			case SDLK_DOWN: vel_y_ = 100; break;
			case SDLK_LEFT: vel_x_ = -100; break;
			case SDLK_RIGHT: vel_x_ = 100; break;
		}
	}
	//A key was released
	else if (e->type == SDL_KEYUP && e->key.repeat == 0)
	{
		switch (e->key.keysym.sym)
		{
			case SDLK_UP: vel_y_ = 0; break;
			case SDLK_DOWN: vel_y_ = 0; break;
			case SDLK_LEFT: vel_x_ = 0; break;
			case SDLK_RIGHT: vel_x_ = 0; break;
		}
	}
}
