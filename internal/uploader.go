package internal

import "context"

type UploadRequest struct {
	Path string
	Tags map[string]string
}

type UploadResponse struct {
}

type IUploader interface {
	Upload(ctx context.Context, request UploadRequest) (UploadResponse, error)
}
