package utils

import "time"

func DurationToString(duration time.Duration) string {
    return duration.String()
}

func FormatDate(t time.Time) string {
    return t.Format(time.RFC3339)
}

func CalculateDuration(start, end time.Time) time.Duration {
    return end.Sub(start)
}