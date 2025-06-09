package tasks

import (
	"math/rand"
	"time"
)

func init() {
	cmd.RegisterTask("iotask", IOTask)
}

func IOTask() (string, error) {
	// Simulate work
	time.Sleep(time.Duration(3+rand.Intn(5)) * time.Minute)
	return "IO Task done", nil
}
