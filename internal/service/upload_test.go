package service

import (
	"context"
	"errors"
	"testing"

	"github.com/problem-01/problem/internal"
)

// mockUploader implements internal.IUploader for testing
type mockUploader struct {
	called bool
	req    internal.UploadRequest
	err    error
}

func (m *mockUploader) Upload(ctx context.Context, req internal.UploadRequest) (internal.UploadResponse, error) {
	m.called = true
	m.req = req
	return internal.UploadResponse{}, m.err
}

func TestUploader_Success(t *testing.T) {
	mock := &mockUploader{}
	uploader := NewUpload(WithUploader(mock))

	err := uploader.Upload(context.Background(), "path/to/file", map[string]string{"env": "test"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !mock.called {
		t.Fatal("expected upload to be called")
	}
	if mock.req.Path != "path/to/file" {
		t.Errorf("expected path to be 'path/to/file', got %s", mock.req.Path)
	}
	if mock.req.Tags["env"] != "test" {
		t.Errorf("expected tag env=test, got %v", mock.req.Tags)
	}
}

func TestUploader_Error(t *testing.T) {
	mock := &mockUploader{err: errors.New("upload failed")}
	uploader := NewUpload(WithUploader(mock))

	err := uploader.Upload(context.Background(), "some/path", nil)

	if err == nil {
		t.Fatal("expected an error but got nil")
	}
	if err.Error() != "upload failed" {
		t.Fatalf("expected 'upload failed', got %v", err)
	}
}
