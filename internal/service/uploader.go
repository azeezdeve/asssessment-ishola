package service

import (
	"context"
	"github.com/problem-01/problem/internal"
)

type IUploader interface {
	Upload(ctx context.Context, path string, tags map[string]string) error
}

type Uploader struct {
	upload internal.IUploader
}

type Config func(Uploader)

func NewUpload(cnf ...Config) Uploader {
	cng := Uploader{}
	for _, c := range cnf {
		c(cng)
	}

	return cng
}

func WithUploader(data internal.IUploader) Config {
	return func(cnf Uploader) {
		cnf.upload = data
	}
}

func (u Uploader) Upload(ctx context.Context, path string, tags map[string]string) error {
	_, err := u.upload.Upload(ctx, internal.UploadRequest{
		Path: path,
		Tags: tags,
	})
	return err
}
