package main

type game struct {
	frames     []frame
	frameIndex uint8
}

func (g *game) initGame() {
	g.frames = make([]frame, 10)
	for indx := range playerGame.frames {
		g.frames[indx].maxRolls = 2
	}
	g.frameIndex = 0
}

func (g *game) calcGameScore() uint16 {
	var gameScore uint16 = 0
	for i := 0; i < len(g.frames); i++ {
		gameScore += g.frames[i].score
	}
	return gameScore
}

func (g *game) isIndexValid() bool {
	if g.frameIndex >= uint8(len(g.frames)) {
		return false
	}
	return true
}
