package ext_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/mocktracer"
)

func TestLogError(t *testing.T) {
	tracer := mocktracer.New()
	span := tracer.StartSpan("my-trace")
	ext.Component.Set(span, "my-awesome-library")
	ext.SamplingPriority.Set(span, 1)
	err := fmt.Errorf("My error")
	ext.LogError(span, err)

	span.Finish()

	rawSpan := tracer.FinishedSpans()[0]
	assert.Equal(t, map[string]interface{}{
		"component": "my-awesome-library",
		"error":     true,
	}, rawSpan.Tags())

	assert.Equal(t, len(rawSpan.Logs()), 1)
	fields := rawSpan.Logs()[0].Fields
	assert.Equal(t, []mocktracer.MockKeyValue{
		{
			Key:         "event",
			ValueKind:   reflect.String,
			ValueString: "error",
		},
		{
			Key:         "message",
			ValueKind:   reflect.String,
			ValueString: err.Error(),
		},
	}, fields)
}
