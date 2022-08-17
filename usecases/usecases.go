package usecases

import (
	"log"
	"net/http"

	"github.com/nelthaarion/artesia/utils/artesiautils"
)

func NewHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			geo := &artesiautils.IPGeo{}
			geo.Load("geo.mmdb")
			ip := artesiautils.ToIP(r.RemoteAddr)
			log.Println(map[string]string{
				"location": geo.GetRecord(ip)["en"],
				"url":      r.URL.Path,
				"method":   r.Method})
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
