package nsqlog

import (
	"io"
	"log"
	"os"

	"github.com/bitly/go-nsq"
)

// Logger - Default consumer Logger.
var Logger *log.Logger

// LogLevel - Default consumer LogLevel.
var LogLevel = nsq.LogLevelDebug

// SetOutput - Changes Logger to new logger with defined output.
func SetOutput(output io.Writer) {
	Logger = log.New(output, "", 0)
}

func init() {
	SetOutput(os.Stdout)
}
