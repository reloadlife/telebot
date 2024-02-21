package log

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var (
	PrefixTag = "TeleBot"
	Prefix    = color.New(color.BgHiWhite, color.FgBlue).Sprintf("[%s]", PrefixTag) + " "
)
var (
	Level = LevelWarn
)

func init() {
	SetLevel(LevelDebug)
	switch strings.ToUpper(os.Getenv("LOG_LEVEL")) {
	case "DEBUG":
		SetLevel(LevelDebug)
	case "INFO":
		SetLevel(LevelInfo)
	case "WARN":
		SetLevel(LevelWarn)
	case "ERROR":
		SetLevel(LevelError)
	case "FATAL":
		SetLevel(LevelFatal)

	default:
		switch strings.ToUpper(os.Getenv("ENV")) {
		case "DEV":
			SetLevel(LevelDebug)
		case "PROD":
			SetLevel(LevelWarn)

		default:
			SetLevel(LevelWarn) // defaults to warning
		}
	}
}

func SetLevel(level int) {
	Level = level
}

func colord(pre, format string, args ...any) (string, string, []any) {
	var clr []color.Attribute

	switch pre {
	case "[DEBUG]":
		clr = append(clr, color.FgHiBlue, color.BgBlack)
	case "[INFO]":
		clr = append(clr, color.FgHiGreen, color.BgBlack)
	case "[WARN]":
		clr = append(clr, color.FgHiYellow, color.BgBlack)
	case "[ERROR]":
		clr = append(clr, color.FgHiRed, color.BgBlack)
	case "[FATAL]":
		clr = append(clr, color.FgHiWhite, color.BgRed)
	default:
		clr = append(clr, color.FgHiWhite, color.BgBlack)
	}
	prefix := color.New(clr...).Sprintf(pre + " ")

	r := regexp.MustCompile(`%[^\s%]*`)
	formatted := r.ReplaceAllString(format, "%v")

	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			args[i] = color.New(color.FgYellow).Sprintf("String(%s)", v)
		case io.Reader:
			stringed := fmt.Sprintf("%v", v)
			args[i] = color.New(color.FgHiYellow).Sprintf("Reader(%s)", stringed)
		case []byte:
			args[i] = color.New(color.FgHiYellow).Sprintf("Bytes(%s)", string(v))
		case int, uint64, int64, uint, uint32, uint16, int32, int16, int8:
			args[i] = color.New(color.FgGreen).Sprintf("Int(%d)", v)
		case float32, float64:
			args[i] = color.New(color.FgHiGreen).Sprintf("Float(%f)", v)
		case bool:
			args[i] = color.New(color.FgHiRed).Sprintf("Bool(%t)", v)
		case nil:
			args[i] = color.New(color.FgHiWhite).Sprintf("Nil(%T)", v)
		default:
			args[i] = color.New(color.FgHiMagenta).Sprintf("%T(%v)", v, v)
		}
	}

	return prefix, formatted, args
}

func Debugf(format string, v ...interface{}) {
	if Level <= LevelDebug {
		prefix, formatted, args := colord("[DEBUG]", format, v...)
		_, _ = fmt.Printf(Prefix+prefix+formatted+"\n", args...)
	}
}
func Infof(format string, v ...interface{}) {
	if Level <= LevelInfo {
		prefix, formatted, args := colord("[INFO]", format, v...)
		_, _ = fmt.Printf(Prefix+prefix+formatted+"\n", args...)
	}
}
func Warnf(format string, v ...interface{}) {
	if Level <= LevelWarn {
		prefix, formatted, args := colord("[WARN]", format, v...)
		_, _ = fmt.Printf(Prefix+prefix+formatted+"\n", args...)
	}
}
func Errorf(format string, v ...interface{}) {
	if Level <= LevelError {
		prefix, formatted, args := colord("[ERROR]", format, v...)
		_, _ = fmt.Printf(Prefix+prefix+formatted+"\n", args...)
	}
}
func Fatalf(format string, v ...interface{}) {
	if Level <= LevelFatal {
		prefix, formatted, args := colord("[FATAL]", format, v...)
		_, _ = fmt.Printf(Prefix+prefix+formatted+"\n", args...)
		panic(fmt.Sprintf(format, v...) + "\n")
	}
}
