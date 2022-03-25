package main

type frame struct {
	maxRolls  uint8
	rollPins  [3]uint8
	rollIndex uint8
	score     uint16
	spare     bool
	strike    bool
	double    bool
	turkey    bool
}

const MAXPINS uint8 = 10
const SPARE uint8 = MAXPINS
const STRIKE uint8 = MAXPINS
const DOUBLE uint8 = 20
const TURKEY uint8 = 30

func (f *frame) initScore() {
	f.score = 0
}

func (f *frame) calcFramesScore() {
	if f.rollIndex > 1 {
		for i := 0; i < int(f.maxRolls); i++ {
			f.score += uint16(f.rollPins[i])
		}
	}
}

func (f *frame) manageIsSpare() {
	if f.maxRolls == 2 && f.score == uint16(SPARE) {
		f.spare = true
	}
}

func (f *frame) isRollValid(fallenPins uint8) bool {
	if fallenPins < 0 || fallenPins > MAXPINS {
		return false
	}
	if f.maxRolls == 2 &&
		f.rollIndex > 0 &&
		fallenPins > (MAXPINS-f.rollPins[0]) {
		return false
	}
	return true
}

func (f *frame) saveFallenPins(fallenPins uint8) {
	f.rollPins[f.rollIndex] = uint8(fallenPins)
}

func (f *frame) manageLastFrameSpare(lastFrame bool) {
	if !lastFrame {
		return
	}
	if f.rollIndex == 2 && (f.rollPins[0]+f.rollPins[1]) == SPARE {
		f.maxRolls = 3
	}
}

func (f *frame) notifyNextRoll() {
	f.rollIndex++
}

func (f *frame) manageStrikes(prevFrame *frame, fallenPins uint8, lastFrame bool) {
	if f.rollIndex == 0 && fallenPins == MAXPINS {
		f.strike = true

		f.maxRolls = 1

		if lastFrame {
			f.maxRolls = 3
		}

		if prevFrame != nil {
			if prevFrame.strike {
				f.double = true
			}
			if prevFrame.double {
				f.turkey = true
			}
		}
	}
}

func (f *frame) impactScoreWithSpare(prevFrame *frame) {
	if prevFrame == nil {
		return
	}
	if prevFrame.spare {
		prevFrame.score += uint16(f.rollPins[0])
	}
}

func (f *frame) impactScoreWithStrike(prevFrame *frame) {
	if prevFrame == nil {
		return
	}
	if prevFrame.strike && !prevFrame.double {
		if f.rollIndex < 2 {
			prevFrame.score = 0
		} else {
			prevFrame.score = uint16(STRIKE) + f.score
		}
	}
}

func (f *frame) impactScoreWithDouble(prevFrame *frame, prevPrevFrame *frame) {
	if prevFrame == nil || prevPrevFrame == nil {
		return
	}
	if prevFrame.double && !f.turkey {
		switch f.rollIndex {
		case 1:
			prevPrevFrame.score = uint16(DOUBLE) + uint16(f.rollPins[0])
		case 2:
			prevPrevFrame.score = uint16(DOUBLE) + uint16(f.rollPins[0])
			prevFrame.score = uint16(STRIKE) + f.score
		}
	}
}

func (f *frame) impactScoreWithTurkey(prevFrame *frame, prevPrevFrame *frame, lastFrame bool) {
	if prevFrame == nil || prevPrevFrame == nil {
		return
	}
	if f.turkey {
		prevPrevFrame.score = uint16(TURKEY)
		if lastFrame && f.rollPins[0] == STRIKE && f.rollPins[1] == STRIKE {
			prevFrame.score = uint16(TURKEY)
		}
	}
}
