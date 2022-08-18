package artesia

import (
	"net/http"

	"github.com/nelthaarion/artesia/usecases"
)

type ArtesiaConfig struct {
	handler http.Handler
}

func NewHandlerGenerator() func(http.Handler) http.Handler {
	return usecases.NewHandler()
}
