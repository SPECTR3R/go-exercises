package cars_test

import (
	"fmt"
	"testing"

	"github.com/SPECTR3R/go-exercises/cars"
)

func TestSuccessRate(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  float64
	}{
		{`0 speed should yield 0%`, 0, 0},
		{`1 - 4 speed should yield 100%`, 3, 1},
		{`5 - 8 speed should yield 90%`, 6, 0.9},
		{`9 - 10 speed should yield 77%`, 9, 0.77},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := cars.SuccessRate(tc.input)
			if s != tc.want {
				t.Errorf("%v success rate.for %d got %v, expected %v", tc.name, tc.input, s, tc.want)
			}
		})
	}
}

func TestCalculateProductionRatePerHour(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  float64
	}{
		{`0 speed should yield 0`, 0, 0},
		{`7 speed should yield 1392.3%`, 7, 1392.3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := cars.CalculateProductionRatePerHour(tc.input)
			if s != tc.want {
				t.Errorf("%v production rate/hour. for %d got %v, expected %v", tc.name, tc.input, s, tc.want)
			}
		})
	}
}

func TestCalculateLimitedProductionRatePerHour(t *testing.T) {
	tt := []struct {
		speed int
		limit float64
		want  float64
	}{
		{0, 500.0, 0.0},
		{2, 1000.0, 442.0},
		{7, 1000.0, 1000.0},
	}

	for _, tc := range tt {
		name := "speed: " + fmt.Sprint(tc.speed)
		t.Run(name, func(t *testing.T) {
			s := cars.CalculateLimitedProductionRatePerHour(tc.speed, tc.limit)
			if s != tc.want {
				t.Errorf("calculate limited production rate per hour for speed %v, got %v, expected %v", tc.speed, s, tc.want)
			}
		})
	}
}
