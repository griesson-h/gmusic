package main
import rl "github.com/gen2brain/raylib-go/raylib"

func renderLoop() {
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable);
	rl.SetTraceLogLevel(rl.LogError);
	rl.InitWindow(800, 600, "gmusic");
	for !rl.WindowShouldClose() {
		rl.BeginDrawing();
		rl.ClearBackground(rl.Gray);
		rl.EndDrawing();
	}
}
