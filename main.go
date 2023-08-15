package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ruleset func(*board, int, int) bool
type board [BOARD_WIDTH][BOARD_HEIGHT]bool

func (b *board) clear() {
	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			b[i][j] = false
		}
	}
}

func (b *board) randomize() {
	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			b[i][j] = (rl.GetRandomValue(0, 1) == 1)
		}
	}
}

func (b *board) randomize_part(p rl.RectangleInt32) {
	for i := p.X; i < p.X+p.Width; i++ {
		for j := p.Y; j < p.Y+p.Height; j++ {
			b[i][j] = (rl.GetRandomValue(0, 1) == 1)
		}
	}
}

func (b *board) copy(dst *board) {
	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			dst[i][j] = b[i][j]
		}
	}
}

func (b *board) next(r ruleset) {
	b.copy(&board_temp)
	for i := 0; i < BOARD_WIDTH; i++ {
		for j := 0; j < BOARD_HEIGHT; j++ {
			board_temp[i][j] = r(b, i, j)
		}
	}
	*b, board_temp = board_temp, *b
}

func (b *board) get_cell(screen_pos rl.Vector2) *bool {
	x := int(screen_pos.X) / CELL_WIDTH
	y := int(screen_pos.Y) / CELL_HEIGHT
	return &b[x][y]
}

func conway(b *board, x, y int) bool {

	alive := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i >= BOARD_WIDTH || y+j < 0 || y+j >= BOARD_HEIGHT {
				continue
			}
			if b[x+i][y+j] {
				alive++
			}
		}
	}
	if b[x][y] {
		return alive == 2 || alive == 3
	} else {
		return alive == 3
	}
}

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 600

	BOARD_WIDTH  = WINDOW_WIDTH / 8
	BOARD_HEIGHT = WINDOW_HEIGHT / 8

	CELL_WIDTH  = WINDOW_WIDTH / BOARD_WIDTH
	CELL_HEIGHT = WINDOW_HEIGHT / BOARD_HEIGHT

	change_speed_cd = float32(0.1)
)

var (
	BOARD             = board{}
	board_temp        = board{}
	fps        uint32 = 30
	ups        uint32 = 3
	dt                = float32(0)

	pause = true

	change_speed_cd_counter = change_speed_cd
)

func change_speed(increase bool) {
	switch increase {
	case true:
		ups += 1
		if ups > 120 {
			ups = 120
		}
	case false:
		ups -= 1
		if ups < 1 {
			ups = 1
		}
	}
	if ups <= 30 && fps > 30 {
		fps = 30
		rl.SetTargetFPS(int32(fps))
	} else if (ups <= 60 && ups > 30) && (fps > 60 || fps <= 30) {
		fps = 60
		rl.SetTargetFPS(int32(fps))
	} else if ups <= 90 && ups > 60 {
		fps = 90
		rl.SetTargetFPS(int32(fps))
	} else if ups <= 120 && ups > 90 {
		fps = 120
		rl.SetTargetFPS(int32(fps))
	}

}

func handle_input() {
	switch rl.GetKeyPressed() {

	case rl.KeyR:
		BOARD.randomize()
	case rl.KeyC:
		BOARD.clear()
	case rl.KeySpace:
		pause = !pause
	case rl.KeyEscape:
		rl.CloseWindow()
	case rl.KeyN:
		if pause {
			update()
		}
	}

	if rl.IsKeyPressed(rl.KeyEqual) || rl.IsKeyPressed(rl.KeyKpAdd) {
		change_speed(true)
	} else if rl.IsKeyDown(rl.KeyEqual) || rl.IsKeyDown(rl.KeyKpAdd) {
		change_speed_cd_counter -= dt
		if change_speed_cd_counter <= 0 {
			change_speed(true)
			change_speed_cd_counter = change_speed_cd
		}
	} else if rl.IsKeyReleased(rl.KeyEqual) || rl.IsKeyReleased(rl.KeyKpAdd) {
		change_speed_cd_counter = change_speed_cd
	}

	if rl.IsKeyPressed(rl.KeyMinus) || rl.IsKeyPressed(rl.KeyKpSubtract) {
		change_speed(false)
	} else if rl.IsKeyDown(rl.KeyMinus) || rl.IsKeyDown(rl.KeyKpSubtract) {
		change_speed_cd_counter -= dt
		if change_speed_cd_counter <= 0 {
			change_speed(false)
			change_speed_cd_counter = change_speed_cd
		}
	} else if rl.IsKeyReleased(rl.KeyMinus) || rl.IsKeyReleased(rl.KeyKpSubtract) {
		change_speed_cd_counter = change_speed_cd
	}

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		*BOARD.get_cell(rl.GetMousePosition()) = true
	}
	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		*BOARD.get_cell(rl.GetMousePosition()) = false
	}
}

func start() {
	BOARD.randomize_part(rl.RectangleInt32{X: 0, Y: 0, Width: 10, Height: 10})
}

func update() {

	BOARD.next(conway)

}

func draw() {
	rl.ClearBackground(rl.GetColor(0x3c3c3cFF))
	for i := int32(0); i < BOARD_WIDTH; i++ {
		for j := int32(0); j < BOARD_HEIGHT; j++ {
			if BOARD[i][j] {
				rl.DrawRectangle(i*CELL_WIDTH, j*CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT, rl.Black)
			} else {
				rl.DrawRectangleLines(i*CELL_WIDTH, j*CELL_HEIGHT, CELL_WIDTH, CELL_HEIGHT, rl.GetColor(0x000000CC))
			}

		}
	}
}

func main() {
	title := "Go of Life"

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, title)
	defer rl.CloseWindow()
	rl.SetTargetFPS(int32(fps))
	start()
	last_update := float32(0)

	for !rl.WindowShouldClose() {
		info := fmt.Sprintf("Current UPS: %v | Current FPS: %v", ups, fps)
		var current_title string
		if pause {
			 current_title = fmt.Sprintf("%s | Paused", title)
		} else {
			 current_title = fmt.Sprintf("%s | %s", title, info)
		}
		rl.SetWindowTitle(current_title)
		update_delay := float32(1) / float32(ups)
		dt = rl.GetFrameTime()
		last_update += dt

		handle_input()

		if last_update >= update_delay {
			if !pause {
				update()
			}
			last_update = 0
		}

		rl.BeginDrawing()
		draw()
		rl.EndDrawing()
	}
}
