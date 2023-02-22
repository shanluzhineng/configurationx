package configurationx

import (
	"fmt"

	jww "github.com/spf13/jwalterweatherman"
)

type myLogger struct{}

func newMyLogger() myLogger {
	// jww.SetLogThreshold(jww.LevelTrace)
	// jww.SetStdoutThreshold(jww.LevelTrace)
	return myLogger{}
}

func (myLogger) Trace(msg string, keyvals ...interface{}) {
	jww.TRACE.Printf(jwwLogMessage(msg, keyvals...))
}

func (myLogger) Debug(msg string, keyvals ...interface{}) {
	jww.DEBUG.Printf(jwwLogMessage(msg, keyvals...))
}

func (myLogger) Info(msg string, keyvals ...interface{}) {
	jww.INFO.Printf(jwwLogMessage(msg, keyvals...))
}

func (myLogger) Warn(msg string, keyvals ...interface{}) {
	jww.WARN.Printf(jwwLogMessage(msg, keyvals...))
}

func (myLogger) Error(msg string, keyvals ...interface{}) {
	jww.ERROR.Printf(jwwLogMessage(msg, keyvals...))
}

func jwwLogMessage(msg string, keyvals ...interface{}) string {
	out := msg

	if len(keyvals) > 0 && len(keyvals)%2 == 1 {
		keyvals = append(keyvals, nil)
	}

	for i := 0; i <= len(keyvals)-2; i += 2 {
		out = fmt.Sprintf("%s %v=%v", out, keyvals[i], keyvals[i+1])
	}

	return out
}
