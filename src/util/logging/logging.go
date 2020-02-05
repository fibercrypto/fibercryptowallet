/*
Package logging provides application logging utilities
*/
package logging

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = NewMasterLogger()
var output = 0

const (
	// logModuleKey is the key used for the module name data entry
	logModuleKey = "_module"
	// logPriorityKey is the log entry key for priority log statements
	logPriorityKey = "_priority"
	// logPriorityCritical is the log entry value for priority log statements
	logPriorityCritical = "CRITICAL"
)

const (
	Debug = iota
	Information
	Warning
	Error
	FatalError
	Panic
)

const (
	Stdout = iota
	Stderr
	None
	File
)

// LevelFromString returns a logrus.Level from a string identifier
func LevelFromString(s string) (logrus.Level, error) {
	switch strings.ToLower(s) {
	case "debug":
		return logrus.DebugLevel, nil
	case "info", "notice":
		return logrus.InfoLevel, nil
	case "warn", "warning":
		return logrus.WarnLevel, nil
	case "error":
		return logrus.ErrorLevel, nil
	case "fatal", "critical":
		return logrus.FatalLevel, nil
	case "panic":
		return logrus.PanicLevel, nil
	default:
		return logrus.DebugLevel, errors.New("Couldn't convert string to log level")
	}
}

// LevelFromEnum returns a logrus.Level from an enum identifier
func LevelFromEnum(s string) (logrus.Level, error) {
	idf, err := strconv.Atoi(s)
	if err != nil {
		return logrus.DebugLevel, errors.New("Couldn't convert enum identifier to log level")
	}

	switch idf {
	case Debug:
		return logrus.DebugLevel, nil
	case Information:
		return logrus.InfoLevel, nil
	case Warning:
		return logrus.WarnLevel, nil
	case Error:
		return logrus.ErrorLevel, nil
	case FatalError:
		return logrus.FatalLevel, nil
	case Panic:
		return logrus.PanicLevel, nil
	default:
		return logrus.DebugLevel, errors.New("Couldn't convert enum identifier to log level")
	}
}

// MustGetLogger returns a package-aware logger from the master logger
func MustGetLogger(module string) *Logger {
	return log.PackageLogger(module)
}

// AddHook adds a hook to the global logger
func AddHook(hook logrus.Hook) {
	log.AddHook(hook)
}

// EnableColors enables colored logging
func EnableColors() {
	log.EnableColors()
}

// DisableColors disables colored logging
func DisableColors() {
	log.DisableColors()
}

// SetLevel sets the logger's minimum log level
func SetLevel(level logrus.Level) {
	log.SetLevel(level)
}

// SetOutputTo sets the logger's output to an io.Writer
func SetOutputTo(w io.Writer) {
	log.Out = w
}

// Disable disables the logger completely
func Disable() {
	log.Out = ioutil.Discard
}

// GetFileToLog get a file with path <dir> for the logger's output
func GetFileToLog(dir string) (io.Writer, error) {
	if output != File {
		return nil, errors.New("Log output isn't a File")
	}

	f, err := os.OpenFile(dir+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0344)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(f)
	//defer f.Close()

	return f, nil
}

// NoWriter is a writer that write nothing, useful when no writing or output is desired.
type NoWriter struct {
	io.Writer
}

func (NoWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

// GetOutputWriter given a option return a writer
func GetOutputWriter(opt string) (io.Writer, error) {
	idf, err := strconv.Atoi(opt)
	if err != nil {
		return nil, errors.New("Couldn't convert enum identifier to log output")
	}

	switch idf {
	case Stdout:
		return os.Stdout, nil
	case Stderr:
		return os.Stderr, nil
	case None:
		return NoWriter{}, nil
	case File:
		output = File
		return nil, nil
	default:
		return nil, errors.New("Couldn't convert enum identifier to log output")
	}
}
