package options

import "github.com/sirupsen/logrus"

type level struct {
	levelFunc func() logrus.Level
}

func (o level) Set(opt *Options) error {
	opt.LevelFunc = o.levelFunc
	return nil
}

func Level(levelFunc func() logrus.Level) Option {
	return level{
		levelFunc: levelFunc,
	}
}
