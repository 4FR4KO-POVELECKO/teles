package teles_test

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/4FR4KO-POVELECKO/teles"
	"github.com/stretchr/testify/assert"
)

func TestLogger_New(t *testing.T) {
	logger := teles.New()
	assert.Equal(t, reflect.TypeOf(logger), reflect.TypeOf(&teles.Logger{}))
}

// var now = time.Now().Format("2006.01.02 15:04:05")
var buf bytes.Buffer
var logger = teles.New()

func readOutput() {
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
}

func TestLogger_NewBot(t *testing.T) {}

func TestLogger_NewDir(t *testing.T) {}

func TestLogger_Log(t *testing.T) {
	readOutput()
	logger.Log(teles.Info, "log")
	t.Log(buf.String())
}

func TestLogger_Trace(t *testing.T) {
	readOutput()
	logger.Trace("trace")
	t.Log(buf.String())
}

func TestLogger_Debug(t *testing.T) {
	readOutput()
	logger.Debug("debug")
	t.Log(buf.String())
}

func TestLogger_Info(t *testing.T) {
	readOutput()
	logger.Info("info")
	t.Log(buf.String())
}

func TestLogger_Warning(t *testing.T) {
	readOutput()
	logger.Warning("warn")
	t.Log(buf.String())
}

func TestLogger_Error(t *testing.T) {
	readOutput()
	logger.Error("error")
	t.Log(buf.String())
}

func TestLogger_Fatal(t *testing.T) {}

func TestLogger_Panic(t *testing.T) {}
