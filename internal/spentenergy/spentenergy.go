package spentenergy

import (
	"errors"
	"time"
)

// константы
const (
	mInKm = 1000 // мин в км
	minInH = 60 // мин в часе
	stepLengthCoefficient = 0.45 // коэф. длины шага
	walkingCaloriesCoefficient = 0.5 // коэф. расчета калорий
)

var (
	errInvalidSteps = errors.New("количество шагов должно быть больше 0")
	errInvalidWeight = errors.New("вес должен быть больше 0")
	errInvalidHeight = errors.New("рост должен быть больше 0")
	errInvalidDuration = errors.New("продолжительность должна быть больше 0")
)

// параметры расчёта калорий
func valid(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return errInvalidSteps
	case weight <= 0:
		return errInvalidWeight
	case height <= 0:
		return errInvalidHeight
	case duration <= 0:
		return errInvalidDuration
	}
	return nil
}

// количество калорий, потраченных при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}

// количество калорий, потраченных при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := valid(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

// средняя скорость в км/ч
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// дистанция в км
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}