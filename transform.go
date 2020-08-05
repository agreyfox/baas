package baas

import (
	"context"
	"time"
)

type TransformService interface {
	Transform(ctx context.Context, f *File, blob *FileBlob, opts TransformationOption) (*FileBlob, error)
	MakeWaterInImage(ctx context.Context, ff *File, blob *FileBlob, text string) (*FileBlob, error)
	VerifyWaterInImage(ctx context.Context, f *File, blob *FileBlob, origf *File, origblob *FileBlob) (*FileBlob, error)
}

type TransformStorage interface {
	Store(ctx context.Context, n *NewTransformation) (*Transformation, error)
}

type Transformation struct {
	ID            string
	ApplicationID string
	FileID        string
	Actions       TransformationOption
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type NewTransformation struct {
	ApplicationID string
	FileID        string
	Actions       TransformationOption
}

type TransformationOption struct {
	// Width
	Width int
	// Height
	Height int
	// Format
	Format string
	// Quality the quality of the JPEG image defaults to 80.
	Quality int
	//  Compression compression for a PNG image defaults to 6.
	Compression int
	// Crop uses lib vips smart crop to crop image
	Crop bool
	// Rotate image rotation angle. Must be a multiple of 90
	Rotate int
	// Flip flips an image
	Flip bool
	// Flop flops an image
	Flop bool
	// Zoom
	Zoom int
	// black and white
	BW bool
	// 是否加水印
	WM bool
	// 水印文字
	WMText string
	//校验WM
	WMV bool
	//Origin
	Origin string
}
