package refasthttp

import (
	"github.com/remicro/api/logging"
	"time"
)

type dummyLogger struct{}

func (d dummyLogger) String(key, value string) logging.Entry {
	return d
}

func (d dummyLogger) Int(key string, value int) logging.Entry {
	return d
}

func (d dummyLogger) Err(err error) logging.Entry {
	return d
}

func (d dummyLogger) Bool(key string, value bool) logging.Entry {
	return d
}

func (d dummyLogger) Time(key string, value time.Time) logging.Entry {
	return d
}

func (d dummyLogger) Duration(key string, duration time.Duration) logging.Entry {
	return d
}

func (d dummyLogger) Float64(key string, value float64) logging.Entry {
	return d
}

func (d dummyLogger) Uint64(key string, value uint64) logging.Entry {
	return d
}

func (d dummyLogger) Logf(message string, args ...interface{}) {

}

func (d dummyLogger) Log(message string) {

}

func (d dummyLogger) Info() logging.Entry {
	return d
}

func (d dummyLogger) Error() logging.Entry {
	return d
}

func (d dummyLogger) Debug() logging.Entry {
	return d
}

func (d dummyLogger) Warn() logging.Entry {
	return d
}

func (d dummyLogger) Critical() logging.Entry {
	return d
}
