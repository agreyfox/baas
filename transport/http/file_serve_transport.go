package http

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/agreyfox/baas"
)

func (s *Server) handleServeFile() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		opts, err := s.parseTransformationOptions(r.URL)
		if err != nil {
			RespondError(w, err, http.StatusInternalServerError, GetReqID(r))
			return
		}

		blob, err := s.FileServeService.Serve(r.Context(), r.URL, opts)
		if err != nil {
			RespondError(w, err, http.StatusInternalServerError, GetReqID(r))
			return
		}
		defer blob.Close()

		RespondServeFile(w, blob)
	})
}

func (s *Server) parseTransformationOptions(u *url.URL) (baas.TransformationOption, error) {
	extension := getFileExtension(u)

	var queryParams serveFileQueryParam
	if err := s.QueryDecoder.Decode(&queryParams, u.Query()); err != nil {
		return baas.TransformationOption{}, fmt.Errorf("failed to decode query params %w", err)
	}

	opts := baas.TransformationOption{
		Width:       queryParams.Width,
		Height:      queryParams.Height,
		Format:      extension,
		Quality:     queryParams.Quality,
		Crop:        queryParams.Crop,
		Flip:        queryParams.Flip,
		Compression: queryParams.Compression,
		Rotate:      queryParams.Rotate,
		Zoom:        queryParams.Zoom,
		BW:          queryParams.BW,
		Flop:        queryParams.Flop,
		WM:          queryParams.WM,
		WMText:      queryParams.WMText,
		WMV:         queryParams.WMV,
		Origin:      queryParams.Origin,
	}

	return opts, nil
}

func getFileExtension(u *url.URL) string {
	return strings.Split(u.Path, ".")[1]
}
