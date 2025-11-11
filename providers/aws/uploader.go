package aws

import (
	"context"
	"errors"
	"fmt"
	"github.com/problem-01/problem/internal"
	"math/rand"
	"os"
	"strings"
	"time"
)

type aws struct {
	AccessKey string
	Secret    string
	Bucket    string
}

func New() internal.IUploader {
	return &aws{
		AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		Secret:    os.Getenv("AWS_SECRET_KEY"),
		Bucket:    os.Getenv("AWS_BUCKET"),
	}
}

// Upload: for uploding to s3 bucket and creating exponential delay
func (a aws) Upload(ctx context.Context, request internal.UploadRequest) (internal.UploadResponse, error) {
	maxDelay := 3
	attempt := 3
	duration := 30 * time.Second
	delay := duration * (1 << attempt)
	if delay > maxDelay {
		return internal.UploadResponse{}, errors.New("max delay exceeded")
	}

	if err := quickUpload(ctx, &MockS3Client{
		AccessKey: a.AccessKey,
		Secret:    a.Secret,
		Bucket:    a.Bucket,
	}, request.Path, request.Tags); err != nil {
		return internal.UploadResponse{}, err
	}
	return internal.UploadResponse{}, nil
}

func quickUpload(ctx context.Context, client *MockS3Client, path string, tags map[string]string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), getFileName(path))

	return client.UploadObject(key, f, tags)
}

func getFileName(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

// ----------------- Mock S3 Client -----------------

type MockS3Client struct {
	AccessKey string
	Secret    string
	Bucket    string
}

func (m *MockS3Client) UploadObject(key string, data []byte, tags map[string]string) error {
	time.Sleep(200 * time.Millisecond)

	// random fail to force candidate to implement retry logic
	if rand.Intn(5) == 0 {
		return errors.New("simulated transient network error")
	}

	fmt.Printf("[MockS3] Uploaded to s3://%s/%s (size=%d) tags=%v\n",
		m.Bucket, key, len(data), tags)

	return nil
}

var _ internal.IUploader = (*aws)(nil)
