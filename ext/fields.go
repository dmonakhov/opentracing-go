package ext

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// These constants define common logFields names recommended for better portability across
// tracing systems and languages/platforms
// More info: https://github.com/opentracing/specification/blob/master/semantic_conventions.md
var (
	LogEvent   = stringLogName("event")
	LogMessage = stringLogName("message")
	LogStack   = stringLogName("stack")

	// Span event represent error
	ErrorEvent = log.String(string(LogEvent), "error")
)

// LogError add error event info for the span
func LogError(span opentracing.Span, err error) {
	Error.Set(span, true)
	span.LogFields(ErrorEvent, log.String(string(LogMessage), err.Error()))
}

// Add log string to to the `span`
func (logName stringLogName) Log(span opentracing.Span, value string) {
	span.LogFields(log.String(string(logName), value))
}

type stringLogName string
