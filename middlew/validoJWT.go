package middlew

import (
	"github.com/hugoing97/twittor/routers"
	"net/http"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), 400)
			return
		}
		next.ServeHTTP(w, r)
	}
}
