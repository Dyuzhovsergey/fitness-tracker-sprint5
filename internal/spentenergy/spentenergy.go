package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lenghtSteps := height * stepLengthCoefficient
	return (float64(steps) * lenghtSteps) / mInKm
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 || steps <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	durationInHour := duration.Hours()
	return distance / durationInHour
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	switch {
	case steps <= 0:
		return 0, errors.New("invalid input: steps must be > 0")
	case weight <= 0:
		return 0, errors.New("invalid input: weight must be > 0")
	case height <= 0:
		return 0, errors.New("invalid input: height must be > 0")
	case duration <= 0:
		return 0, errors.New("invalid input: duration must be > 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	switch {
	case steps <= 0:
		return 0, errors.New("invalid input: steps must be > 0")
	case weight <= 0:
		return 0, errors.New("invalid input: weight must be > 0")
	case height <= 0:
		return 0, errors.New("invalid input: height must be > 0")
	case duration <= 0:
		return 0, errors.New("invalid input: duration must be > 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories * walkingCaloriesCoefficient, nil
}
