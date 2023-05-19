package log

import (
	"errors"
	"github.com/candbright/go-log/options"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInstance(t *testing.T) {
	Instance().Error("Error message")
	Instance().Debug("Debug message")
}

func TestExportedMethod(t *testing.T) {
	Error("Error message")
	Debug("Debug message")
	WithError(errors.New("error happened")).Error("Error message")
}

func TestLevelOpt(t *testing.T) {
	level := logrus.WarnLevel
	err := Init(options.Level(func() logrus.Level {
		return level
	}))
	if err != nil {
		t.Fatal(err)
	}
	Category("WarnLevel").Error("Error message")
	Category("WarnLevel").Warn("Warn message")
	Category("WarnLevel").Info("Info message")
	Category("WarnLevel").Debug("Debug message")

	level = logrus.InfoLevel
	Category("InfoLevel").Error("Error message")
	Category("InfoLevel").Warn("Warn message")
	Category("InfoLevel").Info("Info message")
	Category("InfoLevel").Debug("Debug message")
}

func TestGlobalFieldOpt(t *testing.T) {
	err := Init(options.GlobalField("global", "global value"))
	if err != nil {
		t.Fatal(err)
	}
	Info("message")
}

func TestGlobalFieldSOpt(t *testing.T) {
	err := Init(options.GlobalFields(map[string]interface{}{
		"global field 1": "1",
		"global field 2": "2",
	}))
	if err != nil {
		t.Fatal(err)
	}
	Info("message")
}
