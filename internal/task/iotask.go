package task

import (
	"math/rand"
	"time"
)

type IoTask struct{}

func (t *IoTask) Run() (string, error) {
	time.Sleep(time.Duration(rand.Intn(2)+3) * time.Minute)
	return "Successful result", nil
}
