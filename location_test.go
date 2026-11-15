package bugstack

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// thisFile is the file the tests below call the SDK from. Every capture entry
// point must report this file (and the calling test function), never one of
// the SDK's own files.
func thisFile(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return file
}

func assertCallerLocation(t *testing.T, e *Event, wantFunc string) {
	t.Helper()
	if e == nil {
		t.Fatal("no event was captured")
	}
	if filepath.Base(e.File) != filepath.Base(thisFile(t)) {
		t.Errorf("event.File = %q, want the caller's file %q", e.File, filepath.Base(thisFile(t)))
	}
	if !strings.HasSuffix(e.Function, wantFunc) {
		t.Errorf("event.Function = %q, want the caller %q", e.Function, wantFunc)
	}
}

func initRecording(t *testing.T) **Event {
	t.Helper()
	var got *Event
	Init(Config{
		APIKey:     "bs_test_key",
		DryRun:     true,
		BeforeSend: func(e *Event) *Event { got = e; return e },
	})
	t.Cleanup(func() {
		if c := GetClient(); c != nil {
			c.Shutdown()
		}
	})
	return &got
}

func TestPackageCaptureErrorReportsCaller(t *testing.T) {
	got := initRecording(t)
	CaptureError(errors.New("package-level capture error"))
	assertCallerLocation(t, *got, "TestPackageCaptureErrorReportsCaller")
}

func TestPackageCaptureMessageReportsCaller(t *testing.T) {
	got := initRecording(t)
	CaptureMessage("package-level capture message")
	assertCallerLocation(t, *got, "TestPackageCaptureMessageReportsCaller")
}

func TestClientCaptureErrorReportsCaller(t *testing.T) {
	var got *Event
	c := newTestClient(func(cfg *Config) {
		cfg.BeforeSend = func(e *Event) *Event { got = e; return e }
	})
	defer c.Shutdown()
	c.CaptureError(errors.New("client capture error"))
	assertCallerLocation(t, got, "TestClientCaptureErrorReportsCaller")
}

func TestClientCaptureMessageReportsCaller(t *testing.T) {
	var got *Event
	c := newTestClient(func(cfg *Config) {
		cfg.BeforeSend = func(e *Event) *Event { got = e; return e }
	})
	defer c.Shutdown()
	c.CaptureMessage("client capture message")
	assertCallerLocation(t, got, "TestClientCaptureMessageReportsCaller")
}
