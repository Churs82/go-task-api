package task

import "time"

type Task struct {
    TaskID      string    `json:"task_id"`
    Status      string    `json:"status"`
    CreationDate time.Time `json:"creation_date"`
    Duration    time.Duration `json:"duration"`
}

func (t *Task) SetStatus(status string) {
    t.Status = status
}

func (t *Task) SetDuration(duration time.Duration) {
    t.Duration = duration
}