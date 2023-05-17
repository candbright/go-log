package log

import (
	"github.com/candbright/go-log/options"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInstance(t *testing.T) {
	Instance().Error("Error message")
	Instance().Debug("Debug message")
}

func TestInit(t *testing.T) {
	level := logrus.WarnLevel
	err := Init(options.Level(func() logrus.Level {
		return level
	}))
	if err != nil {
		t.Fatal(err)
	}
	Instance().Category("WarnLevel").Error("Error message")
	Instance().Category("WarnLevel").Warn("Warn message")
	Instance().Category("WarnLevel").Info("Info message")
	Instance().Category("WarnLevel").Debug("Debug message")

	level = logrus.InfoLevel
	Instance().Category("InfoLevel").Error("Error message")
	Instance().Category("InfoLevel").Warn("Warn message")
	Instance().Category("InfoLevel").Info("Info message")
	Instance().Category("InfoLevel").Debug("Debug message")
}
