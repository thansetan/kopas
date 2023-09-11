package helpers

import "time"

const (
	MINUTE = 10 * time.Minute
	HOUR   = time.Hour
	DAY    = 24 * HOUR
	WEEK   = 7 * DAY
)

var ttlMap = map[string]time.Duration{
	"minute": MINUTE,
	"hour":   HOUR,
	"day":    DAY,
	"week":   WEEK,
}

func GetTTL(duration string) time.Duration {
	d, ok := ttlMap[duration]
	if !ok {
		return 0
	}

	return d
}
