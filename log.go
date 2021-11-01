package log

import (
	"log"
	"net/http"
)

type LogLevelType int

const (
	TRACE LogLevelType = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var (
	LogLevel = INFO
)

func (l LogLevelType) String() string {
	switch l {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "INFO"
	}
}

func Println(args ...interface{}) {
	Infoln(args...)
}
func Printf(format string, args ...interface{}) {
	Infof(format, args...)
}

func Fatalln(args ...interface{}) {
	if LogLevel <= FATAL {
		log.Fatal(args...)
	}
}
func Fatalf(format string, args ...interface{}) {
	if LogLevel <= FATAL {
		log.Fatalf(format, args...)
	}
}

func Errorln(args ...interface{}) {
	if LogLevel <= ERROR {
		log.Println(args...)
	}
}
func Errorf(format string, args ...interface{}) {
	if LogLevel <= ERROR {
		log.Printf(format, args...)
	}
}

func Warnln(args ...interface{}) {
	if LogLevel <= WARN {
		log.Println(args...)
	}
}
func Warnf(format string, args ...interface{}) {
	if LogLevel <= WARN {
		log.Printf(format, args...)
	}
}

func Infoln(args ...interface{}) {
	if LogLevel <= INFO {
		log.Println(args...)
	}
}
func Infof(format string, args ...interface{}) {
	if LogLevel <= INFO {
		log.Printf(format, args...)
	}
}

func Debugln(args ...interface{}) {
	if LogLevel <= DEBUG {
		log.Println(args...)
	}
}
func Debugf(format string, args ...interface{}) {
	if LogLevel <= DEBUG {
		log.Printf(format, args...)
	}
}

func Traceln(args ...interface{}) {
	if LogLevel <= TRACE {
		log.Println(args...)
	}
}
func Tracef(format string, args ...interface{}) {
	if LogLevel <= TRACE {
		log.Printf(format, args...)
	}
}

func Error(args ...interface{}) {
	log.Println(args)
}

func Fatal(err error) {
	log.Fatal(err)
}
func Handler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Debugln("http request", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
