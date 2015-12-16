#pragma once

#include <SDL.h>

class Object
{
public:
	Object(int x, int y, int width, int height);

	void render(SDL_Renderer* renderer) const;
	void update(float dt);
	void Object::handle_events(const SDL_Event* e);

private:
	int vel_x_;
	int vel_y_;
	SDL_Rect body_;
};

