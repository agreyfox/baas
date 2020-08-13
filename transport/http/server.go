package http

import (
	"net/http"

	"github.com/agreyfox/baas"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog"
)

type ServerConfig struct {
	BlockService       baas.BlockService
	ApplicationService baas.ApplicationService
	UploadService      baas.UploadService
	FileServeService   baas.FileServeService
	UsageService       baas.UsageService

	QueryDecoder *schema.Decoder

	AuthToken string

	Log zerolog.Logger
}

type Server struct {
	BlockService       baas.BlockService
	ApplicationService baas.ApplicationService
	UploadService      baas.UploadService
	FileServeService   baas.FileServeService
	UsageService       baas.UsageService

	QueryDecoder *schema.Decoder

	AuthToken string

	Log zerolog.Logger

	*mux.Router
}

func NewServer(c *ServerConfig) *Server {
	s := &Server{
		ApplicationService: c.ApplicationService,
		UploadService:      c.UploadService,
		FileServeService:   c.FileServeService,
		UsageService:       c.UsageService,
		BlockService:       c.BlockService,
		QueryDecoder:       c.QueryDecoder,

		AuthToken: c.AuthToken,

		Log: c.Log,

		Router: mux.NewRouter(),
	}

	s.routes()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path

	if r.URL.RawQuery != "" {
		target += "?" + r.URL.RawQuery
	}

	http.Redirect(w, r, target, http.StatusTemporaryRedirect)
}
