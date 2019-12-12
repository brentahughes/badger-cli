package badger

import "fmt"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}
func (l *Logger) Errorf(msg string, args ...interface{}) {
	fmt.Errorf(msg, args)
}
func (l *Logger) Warningf(msg string, args ...interface{}) {

}
func (l *Logger) Infof(msg string, args ...interface{}) {

}
func (l *Logger) Debugf(msg string, args ...interface{}) {

}
