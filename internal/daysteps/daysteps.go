package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("wrong amount of input data")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("failed to convert steps to int: " + err.Error())
	}
	if steps <= 0 {
		return errors.New("steps must be > 0")
	}
	ds.Steps = steps

	d, err := time.ParseDuration(parts[1])
	if err != nil {
		return errors.New("failed to parse duration: " + err.Error())
	}
	if d <= 0 {
		return errors.New("duration must be > 0")
	}
	ds.Duration = d

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	cals, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	res := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cals)
	return res, nil
}
