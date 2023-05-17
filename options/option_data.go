package options

import (
	"io"
)

type filename struct {
	Path string
}

func (o filename) Set(opt *Options) error {
	opt.Path = o.Path
	return nil
}

func Path(path string) Option {
	return filename{
		Path: path,
	}
}

type writer struct {
	Output io.Writer
}

func (o writer) Set(opt *Options) error {
	opt.Output = o.Output
	return nil
}

func Writer(output io.Writer) Option {
	return writer{
		Output: output,
	}
}
