package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/agreyfox/baas/adapter/ipfs"
	"github.com/agreyfox/baas/adapter/s3"
	"github.com/agreyfox/baas/cmd/baasd"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gosimple/slug"
)

func (s *Service) FileBlobStorage(engine, accessKey, secretKey, region, endpoint string) (baas.FileBlobStorage, error) {
	var fileBlobStorage baas.FileBlobStorage
	var err error

	switch engine {
	case baas.StorageEngineAzureBlob:
		return nil, fmt.Errorf("engine not supported yet")
	case baas.StorageEngineIPFS:
		erray := strings.Split(endpoint, ",")
		if len(erray) != 2 {
			fmt.Fprintf(os.Stderr, "Ipfs configuration error, Please check")
			os.Exit(1)
		}
		fileBlobStorage = &ipfs.FileBlobStorage{
			ApiHandle:     erray[0],
			ServiceHandle: erray[1],
		}
		return fileBlobStorage, nil
	default:
		// s3, DO, wasabi, b2

		fileBlobStorage, err = s.s3FileBlobStorage(accessKey, secretKey, region, endpoint)
		if err != nil {
			return nil, fmt.Errorf("cuold not create s3 file blob storage %w", err)
		}
	}

	return fileBlobStorage, nil
}

func (s *Service) createStorageBucket(ctx context.Context, n *baas.NewApplication) error {
	fileBlobStorage, err := s.FileBlobStorage(n.StorageEngine, n.StorageAccessKey, n.StorageSecretKey, n.StorageRegion, n.StorageEndpoint)
	if err != nil {
		return err
	}

	if err := fileBlobStorage.CreateBucket(ctx, n.StorageBucket); err != nil {
		return err
	}

	return nil
}

func (s *Service) s3FileBlobStorage(accessKey, secretKey, region, endpoint string) (*s3.FileBlobStorage, error) {
	// get the user's access keys
	plaintextSecret, err := s.EncryptionService.DecryptString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("could not decrypt storage secret key %w", err)
	}

	// create storage engine configuration
	creds := credentials.NewStaticCredentials(accessKey, string(plaintextSecret), "")
	config := &aws.Config{
		Endpoint:    aws.String(endpoint),
		Region:      aws.String(region),
		Credentials: creds,
		HTTPClient:  &http.Client{Timeout: 5 * time.Minute},
	}

	// create a new s3 session
	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	return &s3.FileBlobStorage{
		AWSSession: sess,
	}, nil
}

func makeStorageEndpoint(engine, region string) string {
	switch engine {
	case baas.StorageEngineDigitalOcean:
		return fmt.Sprintf("%s.digitaloceanspaces.com", region)
	case baas.StorageEngineWasabi:
		return fmt.Sprintf("s3.%s.wasabisys.com", region)
	case baas.StorageEngineB2:
		return fmt.Sprintf("s3.%s.backblazeb2.com", region)
	case baas.StorageEngineIPFS:
		conf := baasd.GetSystemConfig()
		return fmt.Sprintf("%s,%s", conf.IPFS.API, conf.IPFS.Service) //"127.0.0.1:5001,127.0.0.1:10010"
	default:
		return ""
	}
}

func makeStorageBucketName(appname string) string {
	salt := strconv.Itoa(baas.RandInt(1000, 10000))
	sluggedName := appname + "-" + salt

	return slug.Make("baas" + sluggedName)
}
