package teles_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/4FR4KO-POVELECKO/teles"
	"github.com/stretchr/testify/assert"
)

func TestLogger_New(t *testing.T) {
	logger := teles.New()
	assert.Equal(t, reflect.TypeOf(logger), reflect.TypeOf(&teles.Logger{}))
}

var buf bytes.Buffer
var logger = teles.New()

func readOutput() {
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
}

func readLastLine(path string) string {
	file, _ := os.Open(path)
	defer file.Close()

	reader := bufio.NewReader(file)
	lastLineSize := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lastLineSize = len(line)
	}

	fileInfo, _ := os.Stat(path)
	buffer := make([]byte, lastLineSize)
	offset := fileInfo.Size() - int64(lastLineSize+1)
	numRead, _ := file.ReadAt(buffer, offset)

	if offset != 0 {
		buffer = buffer[:numRead]
		return string(buffer)
	}

	return ""
}

func TestLogger_NewBot(t *testing.T) {}

func TestLogger_NewDir(t *testing.T) {
	path := "./log"
	levels := []teles.Level{
		teles.Error,
		teles.Fatal,
	}
	err := logger.NewDir(path, levels)

	assert.NoError(t, err)
	assert.Equal(t, logger.DirPath, path)
	assert.Equal(t, logger.DirLevels, levels)
}

func TestLogger_Log(t *testing.T) {
	readOutput()
	logger.Log(teles.Info, "log")
	t.Log(buf.String())
}

func TestLogger_LogToFile(t *testing.T) {
	var now = time.Now().Format("2006.01.02 15:04:05")
	log := fmt.Sprintf("[%s] ERROR: test error", now)
	path := "./log/error.log"
	logger.LogToFile(path, teles.Error, "test error")
	result, err := os.Stat(path)

	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, os.IsNotExist(err), false)
	assert.Equal(t, result.IsDir(), false)

	lastLine := readLastLine(path)
	fmt.Println("last:", lastLine)
	fmt.Println("log:", log)
	assert.Equal(t, lastLine, log)
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
