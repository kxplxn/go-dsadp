package barrier_pattern

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/headers",
			"http://httpbin.org/user-agent",
		}

		result := captureBarrierOutput(endpoints...)

		containsAcceptEncoding := strings.Contains(result, "Accept-Encoding")
		containsUserAgent := strings.Contains(result, "user-agent")
		if !containsAcceptEncoding || !containsUserAgent {
			t.Fail()
		}

		t.Log(result)
	})

	t.Run("OneIncorrectEndpoint", func(t *testing.T) {
		endpoints := []string{
			"http://malformed-url",
			"http://httpbin.org/user-agent",
		}

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("ShortTimeout", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/headers",
			"http://httpbin.org/user-agent",
		}
		timeoutMilliseconds = 1

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)
	})
}

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()

	os.Stdout = writer

	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out

	return temp
}
