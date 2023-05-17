package options

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Options struct {
	LevelFunc func() logrus.Level
	Formatter logrus.Formatter
	Output    io.Writer
	Path      string
}

func Default() Options {
	return Options{
		LevelFunc: func() logrus.Level {
			return logrus.InfoLevel
		},
		Formatter: &logrus.JSONFormatter{},
		Output:    os.Stdout,
		Path:      "",
	}
}

type Option interface {
	Set(opt *Options) error
}
