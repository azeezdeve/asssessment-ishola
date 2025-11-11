# Problem 1 — Golang AWS Refactoring & Feature Extension

## Goal
You are given an intentionally "quick-and-dirty" single-file implementation that simulates uploading files to S3.  
Your task is to refactor and extend it following sound software design principles.

## Requirements
1. Introduce an `S3Uploader` interface and replace all direct dependencies with DI.
2. Add `context.Context` support to the upload flow.
3. Implement **exponential backoff retry** (max 3 attempts) for transient upload failures.
4. Support object **tags**: `map[string]string`.
5. Remove global variables and replace them with injected configuration.
6. Add at least **one simple unit test** using a mock uploader (no real AWS calls).
7. `go run main.go upload <path>` should still work with the mock client.

## Evaluation Criteria
- Code structure, readability, modularity
- Correct use of interfaces and dependency injection
- Handling of context and cancellation
- Retry correctness (exponential backoff)
- Proper tag passing
- Basic test quality

## How to Run
go run main.go upload ./sample.txt env=staging,team=trading

## Notes
- The provided mock client (`MockS3Client`) must remain usable.
- You may split code into multiple files if desired.