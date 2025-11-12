package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("wrong amount of input data")
	}

	stepsStr := strings.TrimSpace(parts[0])
	kindStr := strings.TrimSpace(parts[1])
	durStr := strings.TrimSpace(parts[2])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return errors.New("failed to convert steps to int: " + err.Error())
	}
	if steps <= 0 {
		return errors.New("steps must be > 0")
	}
	t.Steps = steps
	t.TrainingType = kindStr

	d, err := time.ParseDuration(durStr)
	if err != nil {
		return errors.New("failed to parse duration: " + err.Error())
	}
	if d <= 0 {
		return errors.New("duration must be > 0")
	}
	t.Duration = d

	return nil
}

func (t Training) ActionInfo() (string, error) {
	height := t.Personal.Height
	weight := t.Personal.Weight
	steps := t.Steps
	trainingType := t.TrainingType
	var spentCalories float64
	var err error

	speed := spentenergy.MeanSpeed(steps, height, t.Duration)
	distance := spentenergy.Distance(steps, height)
	switch strings.ToLower(t.TrainingType) {
	case "ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(steps, weight, height, t.Duration)
	case "бег":
		spentCalories, err = spentenergy.RunningSpentCalories(steps, weight, height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, t.Duration.Hours(), distance, speed, spentCalories)
	return result, nil
}
