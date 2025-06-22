package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Init() {
	rl.InitWindow(SCREEN_W, SCREEN_H, TITLE)
	rl.SetTargetFPS(FPS)

	g.running = true
	g.reset = false
}

func (g *Game) Close() {
	rl.CloseWindow()
}

func (g *Game) Loop() {
	for g.running {
		if rl.WindowShouldClose() {
			g.Close()
			g.running = false
			return
		}

		g.update()
		g.draw()
	}
}

func input(g *Game) {
	if !g.reset {
	}

	if g.reset {
	}
}

func (g *Game) update() {
	if (g.reset) {
	}
}

func (g *Game) draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	rl.DrawFPS(10, SCREEN_H-20)

	rl.EndDrawing()
}