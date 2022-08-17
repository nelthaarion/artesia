package artesia

import (
	"net/http"

	"github.com/nelthaarion/artesia/usecases"
)

type ArtesiaConfig struct {
	handler http.Handler
}

func New() func(h http.Handler) http.Handler {
	return usecases.New(http.Handler)
}
