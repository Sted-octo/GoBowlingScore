package main

import "fmt"

var playerGame *game

func main() {
	playerGame = new(game)
	playerGame.initGame()
	fmt.Println("Run the tests, please ;-)")
}

func roll(fallenPins uint8) {
	if !playerGame.isIndexValid() {
		return
	}

	var currentFrame *frame = &playerGame.frames[playerGame.frameIndex]
	var prevFrame *frame = nil
	if playerGame.frameIndex > 0 {
		prevFrame = &playerGame.frames[playerGame.frameIndex-1]
	}

	if !currentFrame.isRollValid(fallenPins) {
		return
	}

	currentFrame.saveFallenPins(fallenPins)

	currentFrame.manageStrikes(prevFrame, fallenPins, playerGame.frameIndex == 9)

	currentFrame.notifyNextRoll()

	currentFrame.manageLastFrameSpare(playerGame.frameIndex == 9)

	if currentFrame.rollIndex >= currentFrame.maxRolls {
		playerGame.frameIndex++
	}
}

func score() uint16 {

	for indx := range playerGame.frames {
		var currentFrame *frame = &playerGame.frames[indx]
		var prevFrame *frame = nil
		var prevPrevFrame *frame = nil
		if indx > 0 {
			prevFrame = &playerGame.frames[indx-1]
		}
		if indx > 1 {
			prevPrevFrame = &playerGame.frames[indx-2]
		}

		currentFrame.initScore()

		currentFrame.calcFramesScore()

		currentFrame.manageIsSpare()

		currentFrame.impactScoreWithSpare(prevFrame)
		currentFrame.impactScoreWithStrike(prevFrame)
		currentFrame.impactScoreWithDouble(prevFrame, prevPrevFrame)
		currentFrame.impactScoreWithTurkey(prevFrame, prevPrevFrame, indx == 9)
	}
	return playerGame.calcGameScore()
}
