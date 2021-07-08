package teles

type Level string

const (
	Panic   Level = "PANIC:"
	Fatal   Level = "FATAL:"
	Error   Level = "ERROR:"
	Warning Level = "WARNING:"
	Info    Level = "INFO:"
	Debug   Level = "DEBUG:"
	Trace   Level = "TRACE:"
)
