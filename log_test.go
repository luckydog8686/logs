package logs

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	fmt.Println("===========")
	Debug("debug")
	Info("info")
	Warn("warn")
	Trace(("trace"))
}