package logger

import (
	"github.com/fatih/color"
)

func Info(message string, a ...interface{}) {
	color.Blue(message, a...)
}

func Error(message string, a ...interface{}) {
	color.Red(message, a...)
}

func Warning(message string, a ...interface{}) {
	color.Yellow(message, a...)
}

func Debug(message string, a ...interface{}) {
	color.Green(message, a...)
}

func Fatal(message string, a ...interface{}) {
	color.Magenta(message, a...)
}
