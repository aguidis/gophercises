package requesthandler

import (
	"net/http"
)

// MapHandler will return an requesthandler.HandlerFunc (which also
// implements requesthandler.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// requesthandler.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(res, req, dest, http.StatusPermanentRedirect)
			return
		}

		fallback.ServeHTTP(res, req)
	}
}
