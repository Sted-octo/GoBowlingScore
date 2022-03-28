package main

import "testing"

func initContext() {
	playerGame = new(game)
	playerGame.initGame()
}

func scoreCompareToExpected(t *testing.T, expected uint16) {
	var currentScore uint16 = score()

	if currentScore != expected {
		t.Errorf("Player score should be %d but was %d", expected, currentScore)
	}
	t.Cleanup(initContext)
}

func Test_score_PaleyrGame_Nil_shouldbe_zero(t *testing.T) {

	var currentScore uint16 = score()

	if currentScore != 0 {
		t.Errorf("Inital score should be 0 but was %d", currentScore)
	}
	t.Cleanup(initContext)
}

func Test_score_initial_values_shouldbe_zero(t *testing.T) {
	initContext()
	var currentScore uint16 = score()

	if currentScore != 0 {
		t.Errorf("Inital score should be 0 but was %d", currentScore)
	}
	t.Cleanup(initContext)
}

func Test_score_1_roll_2_pins_shouldbe_zero(t *testing.T) {
	initContext()
	roll(2)
	scoreCompareToExpected(t, 0)
}

func Test_score_1_roll_12_pins_shouldbe_zero(t *testing.T) {
	initContext()
	roll(12)

	scoreCompareToExpected(t, 0)
}

func Test_score_2_rolls_2_pins_shouldbe_4(t *testing.T) {
	initContext()
	for i := 0; i < 2; i++ {
		roll(2)
	}

	scoreCompareToExpected(t, 4)
}

func Test_score_2_rolls_12_pins_shouldbe_zero(t *testing.T) {
	initContext()
	roll(6)
	roll(6)

	scoreCompareToExpected(t, 0)
}

func Test_score_3_rolls_2_pins_shouldbe_4(t *testing.T) {
	initContext()

	for i := 0; i < 3; i++ {
		roll(2)
	}

	scoreCompareToExpected(t, 4)
}

func Test_score_20_rolls_2_pins_shouldbe_40(t *testing.T) {
	initContext()

	for i := 0; i < 20; i++ {
		roll(2)
	}

	scoreCompareToExpected(t, 40)
}

func Test_score_21_rolls_2_pins_shouldbe_40(t *testing.T) {
	initContext()

	for i := 0; i < 21; i++ {
		roll(2)
	}

	scoreCompareToExpected(t, 40)
}

func Test_score_3_rolls_spare_plus_0_shouldbe_10(t *testing.T) {
	initContext()
	roll(6)
	roll(4)
	roll(0)

	scoreCompareToExpected(t, 10)
}

func Test_score_3_rolls_spare_plus_5_shouldbe_15(t *testing.T) {
	initContext()
	roll(6)
	roll(4)
	roll(5)

	scoreCompareToExpected(t, 15)
}

func Test_score_4_rolls_spare_plus_5_shouldbe_20(t *testing.T) {
	initContext()
	roll(6)
	roll(4)
	roll(5)
	roll(0)

	scoreCompareToExpected(t, 20)
}

func Test_score_5_rolls_2_spares_plus_5_shouldbe_34(t *testing.T) {
	initContext()
	roll(6)
	roll(4)
	roll(9)
	roll(1)
	roll(5)

	scoreCompareToExpected(t, 34)
}

func Test_score_6_rolls_2_spares_plus_5_shouldbe_39(t *testing.T) {
	initContext()
	roll(6)
	roll(4)
	roll(9)
	roll(1)
	roll(5)
	roll(0)

	scoreCompareToExpected(t, 39)
}

func Test_score_21_rolls_10_spares_plus_9_shouldbe_190(t *testing.T) {
	initContext()
	for i := 0; i < 10; i++ {
		roll(9)
		roll(1)
	}
	roll(9)

	scoreCompareToExpected(t, 190)
}

func Test_score_1_roll_strike_10_pins_shouldbe_zero(t *testing.T) {
	initContext()
	roll(10)

	scoreCompareToExpected(t, 0)
}

func Test_score_2_rolls_1_strike_plus_1_shouldbe_zero(t *testing.T) {
	initContext()
	roll(10)
	roll(1)

	scoreCompareToExpected(t, 0)
}

func Test_score_3_rolls_1_strike_1_pin_shouldbe_14(t *testing.T) {
	initContext()
	roll(10)
	roll(1)
	roll(1)

	scoreCompareToExpected(t, 14)
}

func Test_score_3_rolls_1_strike_0_pin_shouldbe_10(t *testing.T) {
	initContext()
	roll(10)
	roll(0)
	roll(0)

	scoreCompareToExpected(t, 10)
}

func Test_score_3_rolls_1_double_plus_1_shouldbe_21(t *testing.T) {
	initContext()
	roll(10)
	roll(10)
	roll(1)

	scoreCompareToExpected(t, 21)
}

func Test_score_4_rolls_1_double_1_pin_shouldbe_35(t *testing.T) {
	initContext()
	roll(10)
	roll(10)
	roll(1)
	roll(1)

	scoreCompareToExpected(t, 35)
}

func Test_score_21_rolls_1_double_2_pins_shouldbe_58(t *testing.T) {
	initContext()
	for i := 0; i < 18; i++ {
		roll(2)
	}
	roll(10)
	roll(10)
	roll(2)

	scoreCompareToExpected(t, 58)
}

func Test_score_3_rolls_1_turkey_shouldbe_30(t *testing.T) {
	initContext()
	for i := 0; i < 3; i++ {
		roll(10)
	}

	scoreCompareToExpected(t, 30)
}

func Test_score_4_rolls_2_turkeys_shouldbe_60(t *testing.T) {
	initContext()
	for i := 0; i < 4; i++ {
		roll(10)
	}

	scoreCompareToExpected(t, 60)
}

func Test_score_12_rolls_world_champion_shouldbe_300(t *testing.T) {
	initContext()
	for i := 0; i < 12; i++ {
		roll(10)
	}

	scoreCompareToExpected(t, 300)
}

func Test_score_12_rolls_11_strikes_plus_3_shouldbe_293(t *testing.T) {
	initContext()
	for i := 0; i < 11; i++ {
		roll(10)
	}
	roll(3)

	scoreCompareToExpected(t, 293)
}

func Test_score_13_rolls_all_strikes_shouldbe_300(t *testing.T) {
	initContext()
	for i := 0; i < 13; i++ {
		roll(10)
	}

	scoreCompareToExpected(t, 300)
}
