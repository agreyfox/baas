package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

type FileBlobStorage struct {
	ApiHandle     string
	ServiceHandle string
	Hash          string
}

func (s *FileBlobStorage) Upload(ctx context.Context, bucket string, b *baas.FileBlob) error {
	fmt.Printf("ipfs buck %s, id %s upload called!", bucket, b.ID)
	//body:=b.Data,
	//	Bucket:      aws.String(bucket),
	//	Key:         aws.String(b.ID),
	//	ContentType: aws.String(b.MIMEValue),

	sh := shell.NewShell(s.ApiHandle)
	/* fmt.Println("Size:", b.Size)
	fmt.Println("Close:", b.MIMEValue)
	fmt.Println("Filename:", b.TempFileName) */
	buf := new(bytes.Buffer)
	buf.ReadFrom(b.Data)
	cid, err := sh.Add(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IPFS error: %s", err)
		//os.Exit(1)
		return err
	}
	fmt.Printf("added %s", cid)
	b.Hash = cid
	return nil
}

func (s *FileBlobStorage) CreateBucket(ctx context.Context, name string) error {
	fmt.Println("Create IPFS bucket ", name)

	return nil
}

func (s *FileBlobStorage) FileBlob(ctx context.Context, bucket, id, tempDir string) (*baas.FileBlob, error) {
	fmt.Println("IPFS fileBlob called:", id)
	buff, err := ioutil.TempFile(tempDir, "ipfs-")
	if err != nil {
		return nil, err
	}
	name := buff.Name()
	buff.Close()
	sh := shell.NewShell(s.ApiHandle)
	err = sh.Get(id, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IPFS Hash %s not read error: %s", id, err)
		return &baas.FileBlob{}, err
	}

	fd, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IPFS generate file  read error: %s", name)
		return &baas.FileBlob{}, err
	}

	f := &baas.FileBlob{
		ID:           "ipfs",
		Data:         ioutil.NopCloser(fd),
		MIMEValue:    "ipfs",
		Size:         1,
		TempFileName: name,
		Hash:         id,
	}

	return f, nil
}
