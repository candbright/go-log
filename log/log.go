package log

import (
	"github.com/candbright/go-log/options"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

func New(opt ...options.Option) (*Logger, error) {
	newLogger := &Logger{
		Logger: logrus.New(),
	}
	err := newLogger.Config(opt...)
	if err != nil {
		return nil, err
	}
	return newLogger, nil
}

type Logger struct {
	*logrus.Logger
	globalFields map[string]interface{}
	levelFunc    func() logrus.Level
}

func (logger *Logger) Config(opt ...options.Option) error {
	o := options.Default()
	var err error
	for _, option := range opt {
		err = option.Set(&o)
		if err != nil {
			return err
		}
	}
	logger.SetFormatter(o.Formatter)
	if o.Path != "" && o.Output == os.Stdout {
		err := os.MkdirAll(path.Dir(o.Path), 0750)
		if err != nil && !os.IsExist(err) {
			return err
		}
		f, err := os.OpenFile(o.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			return err
		}
		logger.SetOutput(f)
	} else {
		logger.SetOutput(o.Output)
	}
	logger.levelFunc = o.LevelFunc
	logger.SetGlobalFields(o.GlobalFields)
	return nil
}

func (logger *Logger) SetGlobalField(key, value string) {
	if logger.globalFields == nil {
		logger.globalFields = make(map[string]interface{})
	}
	logger.globalFields[key] = value
}

func (logger *Logger) SetGlobalFields(fields map[string]interface{}) {
	if logger.globalFields == nil {
		logger.globalFields = make(map[string]interface{})
	}
	for k, v := range fields {
		logger.globalFields[k] = v
	}
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
