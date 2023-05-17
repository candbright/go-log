package options

import "github.com/sirupsen/logrus"

type format struct {
	formatter logrus.Formatter
}

func (o format) Set(opt *Options) error {
	opt.Formatter = o.formatter
	return nil
}

func Format(fmt logrus.Formatter) Option {
	return format{
		formatter: fmt,
	}
}
