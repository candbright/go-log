package log

import (
	"github.com/sirupsen/logrus"
	"go-log/options"
	"os"
)

var logger *Logger

func Init(opt ...options.Option) error {
	o := options.Default()
	var err error
	for _, option := range opt {
		err = option.Set(&o)
		if err != nil {
			return err
		}
	}
	if logger == nil {
		logger = &Logger{
			Logger: &logrus.Logger{},
		}
	}
	logger.SetFormatter(o.Formatter)
	if o.Path != "" && o.Output == os.Stdout {
		f, err := os.OpenFile(o.Path, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logger.SetOutput(f)
	} else {
		logger.SetOutput(o.Output)
	}
	logger.levelFunc = o.LevelFunc
	return nil
}

func Instance() *Logger {
	if logger == nil {
		return nil
	}
	logger.SetLevel(logger.levelFunc())
	return logger
}

type Logger struct {
	*logrus.Logger
	levelFunc func() logrus.Level
}

func (logger *Logger) Category(category string) *logrus.Entry {
	return Instance().Logger.WithFields(logrus.Fields{"category": category})
}
