package logger

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type LogFormatterImpl struct {
	Logger Common
}

func (l *LogFormatterImpl) NewLogEntry(r *http.Request) middleware.LogEntry {
	from := r.RemoteAddr
	method := r.Method

	uri := r.RequestURI

	entry := LogEntryImpl{
		log:     l.Logger,
		request: r,
		buf:     &bytes.Buffer{},
	}

	_, err := fmt.Fprintf(entry.buf, "[%s]  > %s - from: %s - ", method, uri, from)
	if err != nil {
		l.Logger.Error(err, "Failed converting Logger entry")
	}

	return entry
}

type LogEntryImpl struct {
	log     Common
	request *http.Request
	buf     *bytes.Buffer
}

func (l LogEntryImpl) Write(status, bytes int, _ http.Header, elapsed time.Duration, _ interface{}) {
	if _, err := fmt.Fprintf(l.buf, "%db in %s - %d", bytes, elapsed, status); err != nil {
		l.log.Error(err, "Failed converting Logger entry")
	}

	if status >= 500 {
		l.log.Error(errors.New("Http error"), l.buf.String())
	} else {
		l.log.Info(l.buf.String())
	}
}

func (l LogEntryImpl) Panic(v interface{}, stack []byte) {
	l.log.Errorf(errors.New("Panic!"), "Error: %v, Stack: %s", v, string(stack))
}
