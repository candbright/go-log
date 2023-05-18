package log

import (
	"github.com/candbright/go-log/options"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	*logrus.Logger
	levelFunc func() logrus.Level
}

func (logger *Logger) Category(category string) *logrus.Entry {
	return logger.Logger.WithFields(logrus.Fields{"category": category})
}

func New(opt ...options.Option) (*Logger, error) {
	o := options.Default()
	var err error
	for _, option := range opt {
		err = option.Set(&o)
		if err != nil {
			return nil, err
		}
	}
	newLogger := &Logger{
		Logger: logrus.New(),
	}
	newLogger.SetFormatter(o.Formatter)
	if o.Path != "" && o.Output == os.Stdout {
		f, err := os.OpenFile(o.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			return nil, err
		}
		newLogger.SetOutput(f)
	} else {
		newLogger.SetOutput(o.Output)
	}
	newLogger.levelFunc = o.LevelFunc
	return newLogger, nil
}
