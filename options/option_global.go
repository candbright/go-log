package options

type globals struct {
	GlobalFields map[string]interface{}
}

func (o globals) Set(opt *Options) error {
	if opt.GlobalFields == nil {
		opt.GlobalFields = make(map[string]interface{})
	}
	for k, v := range o.GlobalFields {
		opt.GlobalFields[k] = v
	}
	return nil
}

func GlobalFields(fields map[string]interface{}) Option {
	return globals{
		GlobalFields: fields,
	}
}

type global struct {
	Key   string
	Value interface{}
}

func (o global) Set(opt *Options) error {
	if opt.GlobalFields == nil {
		opt.GlobalFields = make(map[string]interface{})
	}
	opt.GlobalFields[o.Key] = o.Value
	return nil
}

func GlobalField(key string, value interface{}) Option {
	return global{
		Key:   key,
		Value: value,
	}
}
