package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"strings"
	"time"
)

func rotate(maxAge int, rotationTime time.Duration,logPath string, format string) (r *rotatelogs.RotateLogs, err error) {
	tmp := strings.Split(logPath, ".")
	if tmp[len(tmp)-1] == "log" {
		tmp = tmp[:len(tmp)-1]
	}

	logPath = strings.Join(tmp, ".") + "-" + format + ".log"
	r, err = rotatelogs.New(
		logPath,
		rotatelogs.WithMaxAge(time.Duration(int64(24*time.Hour) * int64(maxAge))),
		rotatelogs.WithRotationTime(rotationTime),
	)
	return
}
