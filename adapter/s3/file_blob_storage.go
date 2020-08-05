package s3

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/agreyfox/baas"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws/session"
)

type FileBlobStorage struct {
	AWSSession *session.Session
}

func (s *FileBlobStorage) Upload(ctx context.Context, bucket string, b *baas.FileBlob) error {
	uploader := s3manager.NewUploader(s.AWSSession)
	if _, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Body:        b.Data,
		Bucket:      aws.String(bucket),
		Key:         aws.String(b.ID),
		ContentType: aws.String(b.MIMEValue),
	}); err != nil {
		if err := handleAWSErr(err); err != nil {
			return err
		}
		return fmt.Errorf("could not upload file to s3 compatiable storage %w", err)
	}

	return nil
}

func (s *FileBlobStorage) CreateBucket(ctx context.Context, name string) error {
	svc := s3.New(s.AWSSession)
	input := &s3.CreateBucketInput{
		Bucket: aws.String(name),
	}

	_, err := svc.CreateBucketWithContext(ctx, input)
	if err != nil {
		if err := handleAWSErr(err); err != nil {
			return err
		}
		return fmt.Errorf("could not create bucket %w", err)
	}

	if err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(name),
	}); err != nil {
		if err := handleAWSErr(err); err != nil {
			return err
		}
		return err
	}

	return nil
}

func (s *FileBlobStorage) FileBlob(ctx context.Context, bucket, id, tempDir string) (*baas.FileBlob, error) {
	downloader := s3manager.NewDownloader(s.AWSSession)

	svc := s3.New(s.AWSSession)

	resp, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(id),
	})
	if err != nil {
		return nil, fmt.Errorf("could not get object meta data %w", err)
	}

	buff, err := ioutil.TempFile(tempDir, "s3-")
	if err != nil {
		return nil, err
	}

	if _, err := downloader.DownloadWithContext(ctx, buff, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(id),
	}); err != nil {
		if err := handleAWSErr(err); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("could not get file from s3 compatible storage %w", err)
	}

	f := &baas.FileBlob{
		ID:           id,
		Data:         buff,
		MIMEValue:    *resp.ContentType,
		Size:         *resp.ContentLength,
		TempFileName: buff.Name(),
	}

	return f, nil
}

func handleAWSErr(err error) error {
	var awsErr awserr.Error
	if errors.As(err, &awsErr) {
		switch awsErr.Code() {
		case s3.ErrCodeNoSuchBucket:
			return baas.ErrBucketNotFound
		case s3.ErrCodeNoSuchKey:
			return baas.ErrFileNotFound
		case s3.ErrCodeNoSuchUpload:
			return baas.ErrFileNotFound
		case "InvalidAccessKeyId":
			return baas.ErrInvalidStorageKey
		case "AccessDenied":
			return baas.ErrInvalidStorageKey
		}

		if awsErr.Message() == "Access Denied" {
			return baas.ErrInvalidStorageKey
		}
	}

	return nil
}
