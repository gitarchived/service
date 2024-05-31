package s3

import (
	"context"
	"os"

	"github.com/minio/minio-go/v7"
)

func (s *S3) Upload(file string) error {
	var ctx = context.Background()

	_, err := s.FPutObject(
		ctx,
		os.Getenv("STORAGE_BUCKET"),
		file,
		file,
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		},
	)

	return err
}
