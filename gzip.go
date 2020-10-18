package gziphandler

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

// Optional Pool
// var (
// 	pool = sync.Pool{
// 		New: func() interface{} {
// 			w,_ := gzip.NewWriterLevel(nil,1)
// 			return &gzipResponseWriter{
// 				w:w,
// 			}
// 		}
// 	}
// )

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}

// Gzipler is the middleware
func Gzipler(h http.Handler, level int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz, _ := gzip.NewWriterLevel(w, level)
		defer gz.Close()
		gw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		h.ServeHTTP(gw, r)
	})
}
