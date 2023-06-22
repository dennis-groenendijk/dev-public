package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"go-syntax/bits-and-bytes/pkg/path"
)

// explain: go.mod, packages & import

// define global game variables
// explain: single/literal, list, short assignment, type inference & function output
var (
	avatarPosX   int32
	avatarPosY   int32
	timer        = 60 * 60 * 1 // 1 minute
	frameCounter int
	scoreCounter int
)

const screenWidth, screenHeight = int32(1280), int32(720)

type Bit struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	color  rl.Color
}

// explain: composition/embedding
type Player struct {
	name      string
	highscore int
	avatar    Avatar
}

type Avatar struct {
	posX int
	posY int
}

// explain: functions vs methods
// explain: private vs public in Go
func spawnPlayer(x, y int) (bool, error) {
	// todo: something
	return false, nil
}

// explain: method is a function with receiver argument
func (p *Player) spawn(x, y int) (bool, error) {
	// todo: something
	return false, nil
}

func main() {
	// demo.Demo(1) // demonstrate calling functions from another package

	// game settings
	rl.InitWindow(screenWidth, screenHeight, "Some dude presents BITS AND BYTES")
	rl.SetTargetFPS(60)

	// define avatar
	avatarDown := rl.LoadImage(path.Root + "/bit-and-bytes/assets/flappy-down.png")
	avatarUp := rl.LoadImage(path.Root + "/bit-and-bytes/assets/flappy-up.png")
	avatar := rl.LoadTextureFromImage(avatarUp)

	// define data bits (= targets)
	rand.Seed(time.Now().UnixNano())

	var eightBits []Bit
	var posBit = int32(rand.Intn(720-2+1) - 2)
	bitRedS := Bit{screenWidth, posBit, 25, 25, rl.Red}
	bitRedM := Bit{screenWidth, posBit, 65, 35, rl.Red}
	bitRedL := Bit{screenWidth, posBit, 185, 25, rl.Red}

	bitGreenS := Bit{screenWidth, posBit, 25, 25, rl.Green}
	bitGreenM := Bit{screenWidth, posBit, 65, 35, rl.Green}
	bitGreenL := Bit{screenWidth, posBit, 185, 25, rl.Green}

	bitBlueS := Bit{screenWidth, posBit, 25, 25, rl.Blue}
	bitBlueM := Bit{screenWidth, posBit, 65, 35, rl.Blue}
	eightBits = append(eightBits, bitRedS, bitRedM, bitRedL, bitGreenS, bitGreenM, bitGreenL, bitBlueS, bitBlueM)

	// define multiple game screens
	// explain: maps
	gameScreen := make(map[int]string, 0)
	gameScreen[0] = "START"
	gameScreen[1] = "GAME"
	gameScreen[2] = "GAMEOVER"

	currentScreen := "START"

	// runtime
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// game screen manager
		// explain: Go only runs selected case, no need for break statements
		// explain: cases don't have to be constants and values don't have to be integers
		// explain: switch condition/outcome vs no condition/case is condition
		switch currentScreen {
		case "START":
			// position avatar
			avatarPosX = screenWidth/2 - avatar.Width/2
			avatarPosY = screenHeight/2 - avatar.Height/2 - 40

			// (re)set counters
			frameCounter = timer
			scoreCounter = 0

			// confirm selection
			if rl.IsKeyPressed(rl.KeySpace) {
				currentScreen = "GAME"
			}

			// start screen
			flappy := rl.LoadImage(path.Root + "/bit-and-bytes/assets/flappy-up.png")
			splash := rl.LoadTextureFromImage(flappy)
			rl.DrawTexture(splash, 600, 400, rl.White)

			rl.DrawText("Press or Hold Space to push Flappy up. Release to let Flappy go down", 240, 180, 20, rl.LightGray)
			rl.DrawText("Score as many points as possible by collecting bits and bytes of data, but don't fly off the screen!!", 120, 220, 20, rl.LightGray)
			rl.DrawText("Press Space to start", 400, 300, 40, rl.LightGray)
		case "GAME":
			// start timer
			frameCounter--

			// spawn avatar
			rl.DrawTexture(avatar, avatarPosX, avatarPosY, rl.White)

			rl.DrawText("Current Score: "+strconv.Itoa(scoreCounter), 0, 0, 20, rl.LightGray)

			// game stages/difficulty
			if frameCounter > timer-timer/3 {
				rl.DrawText("Time: "+strconv.Itoa(frameCounter), screenWidth-105, 0, 20, rl.LightGray)
			} else if frameCounter > timer-2*(timer/3) {
				rl.DrawText("Time: "+strconv.Itoa(frameCounter), screenWidth-105, 0, 20, rl.Orange)
			} else {
				rl.DrawText("Time: "+strconv.Itoa(frameCounter), screenWidth-105, 0, 20, rl.Red)
			}

			// control avatar
			if rl.IsKeyDown(rl.KeySpace) {
				avatar = rl.LoadTextureFromImage(avatarUp)
				avatarPosY -= 15
			} else {
				avatar = rl.LoadTextureFromImage(avatarDown)
				avatarPosY += 10
			}

			// load a byte of data
			Byte := eightBits

			// control bits and scoring
			// explain: for loops (while, infinite loops)
			// explain: demo identifiers
			for i, bit := range Byte {
				rl.DrawRectangle(bit.posX, bit.posY, bit.width, bit.height, bit.color)
				colliderBit := rl.NewRectangle(float32(bit.posX), float32(bit.posY), float32(bit.width), float32(bit.height))
				colliderScraper := rl.NewRectangle(float32(avatarPosX), float32(avatarPosY), float32(34), float32(24))

				// speed of bits
				if frameCounter > timer-timer/3 {
					Byte[i].posX = Byte[i].posX - 15 // starting speed
				} else if frameCounter > timer-2*(timer/3) {
					Byte[i].posX = Byte[i].posX - 25 // increased speed
				} else {
					Byte[i].posX = Byte[i].posX - 55 // insane speed
				}

				// bit missed
				if bit.posX < 0 {
					Byte[i].posX = screenWidth
					Byte[i].posY = int32(rand.Intn(int(screenHeight-2+1)) - 2)
				}

				// bit hit
				if rl.CheckCollisionRecs(colliderScraper, colliderBit) {
					Byte[i].posX = screenWidth
					Byte[i].posY = int32(rand.Intn(int(screenHeight-2+1)) - 2)
					scoreCounter++
				}
			}

			// game over
			if avatarPosY < 0 || avatarPosY > screenHeight || frameCounter == 0 {
				Byte = nil
				frameCounter = 0
				rl.UnloadTexture(avatar)
				currentScreen = "GAMEOVER"
			}
		case "GAMEOVER":
			rl.DrawText("GAME OVER", screenWidth/2-305, 80, 100, rl.Red)
			rl.DrawText("Final Score: "+strconv.Itoa(scoreCounter), screenWidth/2-150, 180, 40, rl.LightGray)

			// to continue
			rl.DrawText("Press Space to continue", 500, 260, 20, rl.LightGray)
			if rl.IsKeyPressed(rl.KeySpace) {
				currentScreen = "START"
			}

			// to quit game
			rl.DrawText("Press Esc to quit", 500, 340, 20, rl.LightGray)
		}
		rl.EndDrawing()
		time.Sleep(50000000)
	}
	// shut down game
	rl.CloseWindow()
}
