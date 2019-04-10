package log

import (
	"github.com/op/go-logging"
	"os"
)

var (
	Logger *logging.Logger
	format logging.Formatter
)

func InitLogger(logSuffix string) {
	Logger = logging.MustGetLogger("cms-admin-api")
	format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfile} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)
	f, _ := os.Create("/tmp/cms-admin-api.log")
	backend1 := logging.NewLogBackend(f, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	logging.SetBackend(backend1Formatter, backend2Formatter)
}
