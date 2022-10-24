package logger

//go:generate ../../../bin/mockery --name=Logger --case underscore

import (
	"fmt"
)

type Logger interface {
	Log(v ...interface{})
	Panic(v ...interface{})
	Print(v interface{})
}

type logger struct {
}

func New() Logger {
	return &logger{}
}

func (l *logger) Log(v ...interface{}) {
	fmt.Println(v...)
}

func (l *logger) Print(v interface{}) {
	//aYAML, err := yaml.Marshal(v)
	//if err != nil {
	//	l.Log(v)
	//} else {
	//fmt.Printf("YAML Print - \n%s\n", string(aYAML))
	//}
}

func (l *logger) Panic(v ...interface{}) {
	fmt.Println(v...)
	panic(struct{}{})
}
