# Changelog

All notable changes to this project will be documented in this file.

## [1.2.0] - 2026-10-28

### Fixed

- The default endpoint is now `https://dashboard.bugstack.ai/api/ingest/sdk`.
  It was `https://api.bugstack.dev/api/capture`, which is not bugstack's
  ingest address, so events sent with the default setting did not reach
  bugstack. Apps that set `Config.Endpoint` themselves are unaffected.
- The README link and security address now point at bugstack.ai.

## [1.0.0] - 2026-02-13

### Added

- Core SDK with `bugstack.Init()`, `bugstack.CaptureError()`, `bugstack.CaptureMessage()`
- Panic recovery with `bugstack.Recover()`
- Background goroutine transport with buffered channel and retry with exponential backoff
- SHA-256 error fingerprinting and client-side deduplication
- `BeforeSend` hook for event inspection/modification/filtering
- `IgnoredErrors` for skipping errors by message substring
- `DryRun` mode for transparent debugging
- `Enabled` kill switch
- net/http middleware with panic recovery
- Gin middleware adapter
- Echo middleware adapter
- Zero external dependencies
