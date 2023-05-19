package log

import (
	"github.com/candbright/go-log/options"
	"github.com/sirupsen/logrus"
	"os"
)

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
	newLogger.GlobalFields = o.GlobalFields
	return newLogger, nil
}

type Logger struct {
	*logrus.Logger
	GlobalFields map[string]interface{}
	levelFunc    func() logrus.Level
}

func (logger *Logger) Category(category string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{"category": category})
}

type Entry struct {
	*logrus.Entry
}

func (entry *Entry) Category(category string) *logrus.Entry {
	return entry.WithFields(logrus.Fields{"category": category})
}
