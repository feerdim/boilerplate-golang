package log

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/getsentry/sentry-go"
)

func parseInt(i int, s string) int {
	if s == "" {
		return i
	}

	o, e := strconv.Atoi(s)
	if e != nil || o <= 0 {
		return i
	}

	return o
}

func ParseJSON(data interface{}) string {
	JSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(JSON)
}

func ParsePrettyJSON(data interface{}) string {
	JSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(JSON)
}

func generateMessage(msg string, fields []interface{}) string {
	if len(fields) == 0 {
		return msg
	}

	var json string

	if len(fields) > 1 {
		for i := 1; i < len(fields); i++ {
			json = fmt.Sprintf("%s\n%s : %s", json, fields[i-1], ParseJSON(fields[i]))
			i++
		}
	}

	return fmt.Sprintf("%s%s", msg, json)
}

func (l *Logger) sendSentry(msg string, level int) {
	l.sentry.Level = parseLevelSentry(level)
	l.sentry.Message = msg

	_ = sentry.CaptureEvent(l.sentry)
}

func (l *Logger) sendExceptionSentry(err error, msg string, level int, pc uintptr, file string, line int) {
	l.sentry.SetException(err, 1)
	l.sentry.Level = parseLevelSentry(level)
	l.sentry.Message = msg

	l.sentry.Exception[0].Stacktrace.Frames[0].Function = runtime.FuncForPC(pc).Name()
	l.sentry.Exception[0].Stacktrace.Frames[0].Filename = filepath.Base(file)
	l.sentry.Exception[0].Stacktrace.Frames[0].AbsPath = file
	l.sentry.Exception[0].Stacktrace.Frames[0].Lineno = line

	_ = sentry.CaptureEvent(l.sentry)
}
