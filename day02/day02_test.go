package day02_test

import (
	"testing"

	"github.com/m0sh1x2/advent_of_code_2022/day02"
)

func TestDay02Part01(t *testing.T) {

	t.Run("DefaultTest", func(t *testing.T) {
		inputs := []string{
			"A Y",
			"B X",
			"C Z",
		}
		got := day02.Day02Part01(inputs)
		want := 15

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, inputs)
		}
	})

	t.Run("Draw", func(t *testing.T) {
		inputs := []string{
			"A A",
		}
		got := day02.Day02Part01(inputs)
		want := 4

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, inputs)
		}
	})

	t.Run("Win", func(t *testing.T) {
		inputs := []string{
			"A B",
		}
		got := day02.Day02Part01(inputs)
		want := 8

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, inputs)
		}
	})

	t.Run("Loose", func(t *testing.T) {
		inputs := []string{
			"A C",
		}
		got := day02.Day02Part01(inputs)
		want := 3

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, inputs)
		}
	})

}

func TestDay02Part02(t *testing.T) {

	t.Run("DefaultTest", func(t *testing.T) {
		inputs := []string{
			"A Y",
			"B X",
			"C Z",
		}
		got := day02.Day02Part02(inputs)
		want := 12

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, inputs)
		}
	})

	t.Run("Draw", func(t *testing.T) {
		inputs := []string{
			"A Y",
		}
		got := day02.Day02Part02(inputs)
		want := 4

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, inputs)
		}
	})

	t.Run("Loose", func(t *testing.T) {
		inputs := []string{
			"B X",
		}
		got := day02.Day02Part02(inputs)
		want := 1

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, inputs)
		}
	})

	t.Run("Win", func(t *testing.T) {
		inputs := []string{
			"C Z",
		}
		got := day02.Day02Part02(inputs)
		want := 7

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, inputs)
		}
	})

}
