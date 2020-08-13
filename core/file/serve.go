package file

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/agreyfox/baas"
	"github.com/rs/zerolog"
)

type ServeService struct {
	FileStorage baas.FileStorage

	UsageService       baas.UsageService
	ApplicationService baas.ApplicationService
	//	TransformService   baas.TransformService

	FullFileDir string

	Log zerolog.Logger
}

func (s *ServeService) Serve(ctx context.Context, u *url.URL, opts baas.TransformationOption) (*baas.FileBlob, error) {
	fileBlobID := getFileBlobID(u.Path)

	file, err := s.FileStorage.FileByFileBlobID(ctx, fileBlobID)
	if err != nil {
		return nil, err
	}

	app, err := s.ApplicationService.Application(ctx, file.ApplicationID)
	if err != nil {
		return nil, err
	}

	fileBlobStorage, err := s.ApplicationService.FileBlobStorage(app.StorageEngine, app.StorageAccessKey, app.StorageSecretKey, app.StorageRegion, app.StorageEndpoint)
	if err != nil {
		return nil, err
	}
	var id string
	if len(file.Hash) > 0 {
		id = file.Hash
	} else {
		id = fileBlobID
	}

	fileBlob, err := fileBlobStorage.FileBlob(ctx, app.StorageBucket, id, s.FullFileDir)
	if err != nil {
		return nil, fmt.Errorf("could not get file blob %w", err)
	}
	if len(file.Hash) > 0 { //fixed for ipfs
		fileBlob.Size = file.Size
		fileBlob.MIMEValue = file.MIMEValue
	}
	/*	if opts.WM { // should to do the blind watermark
			s.Log.Printf("Doing the blind wm job")
			//return fileBlob, nil
			//s.Log.Printf("%v,%v", file, fileBlob)
			f, err := s.TransformService.MakeWaterInImage(ctx, file, fileBlob, opts.WMText)
			if err != nil {
				return &baas.FileBlob{}, err
			} else {
				return f, nil
			}
		}
		if opts.WMV {
			s.Log.Printf("Doing the verify Blind WaterMark job")
			// Now we have check image, then we need find the origin imagefile
			origfileBlobID := getFileBlobID(opts.Origin)

			origfile, err := s.FileStorage.FileByFileBlobID(ctx, origfileBlobID)
			if err != nil {
				return nil, err
			}

			origapp, err := s.ApplicationService.Application(ctx, origfile.ApplicationID)
			if err != nil {
				return nil, err
			}

			origfileBlobStorage, err := s.ApplicationService.FileBlobStorage(origapp.StorageEngine, origapp.StorageAccessKey, origapp.StorageSecretKey, origapp.StorageRegion, origapp.StorageEndpoint)
			if err != nil {
				return nil, err
			}
			var id string
			if len(origfile.Hash) > 0 {
				id = origfile.Hash
			} else {
				id = origfileBlobID
			}

			origfileBlob, err := origfileBlobStorage.FileBlob(ctx, origapp.StorageBucket, id, s.FullFileDir)
			if err != nil {
				return nil, fmt.Errorf("could not get origin file blob %w", err)
			}
			if len(origfile.Hash) > 0 { //fixed for ipfs
				origfileBlob.Size = file.Size
				origfileBlob.MIMEValue = file.MIMEValue
			}
			// found.check watermrk
			f, err := s.TransformService.VerifyWaterInImage(ctx, file, fileBlob, origfile, origfileBlob)
			// get the result and return
			if err != nil {
				return &baas.FileBlob{}, err
			} else {
				return f, nil
			}
		}
	*/
	defer fileBlob.Close()
	//if !shouldTransform(file, opts) {
	updatedUsages := &baas.UpdateUsage{
		ApplicationID: file.ApplicationID,
		Bandwidth:     fileBlob.Size,
	}

	if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
		// should I fail the request
		s.Log.Error().Err(err).Msg("failed to update usage")
	}

	return fileBlob, nil
	//}

	// we can close the original fileBlob since we will be transforming it and generating a new one.
	// the returned blob gets closed by the parent of this function that still needs the blob around.

	/*
			transformedBlob, err := s.TransformService.Transform(ctx, file, fileBlob, opts)
			if err != nil {
				return nil, err
			}

		updatedUsages := &baas.UpdateUsage{
			ApplicationID: file.ApplicationID,
			Bandwidth:     transformedBlob.Size,
		}

		if err := s.UsageService.Update(ctx, updatedUsages); err != nil {
			// should I fail the request
			s.Log.Error().Err(err).Msg("failed to update usage")
		}

		return transformedBlob, nil*/
}

func shouldTransform(file *baas.File, opts baas.TransformationOption) bool {
	return opts.Width > 0 ||
		opts.Height > 0 ||
		opts.Format != file.Extension
}

func getFileBlobID(urlPath string) string {
	path := strings.TrimPrefix(urlPath, "/")
	return strings.Split(path, ".")[0]
}
