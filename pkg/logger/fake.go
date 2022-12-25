package logger

type Fake struct {
}

func NewFake() *Fake {
	return &Fake{}
}

func (f *Fake) Named(string) Logger {
	return f
}

func (f *Fake) With(...interface{}) Logger {
	return f
}

func (*Fake) Info(...interface{}) {

}

func (*Fake) Infof(string, ...interface{}) {

}

func (*Fake) Infow(string, ...interface{}) {

}

func (*Fake) Warn(...interface{}) {

}

func (*Fake) Warnf(string, ...interface{}) {

}

func (*Fake) Error(...interface{}) {

}

func (*Fake) Errorf(string, ...interface{}) {

}

func (*Fake) Fatal(...interface{}) {

}

func (*Fake) Fatalf(string, ...interface{}) {

}
