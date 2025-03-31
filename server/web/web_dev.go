//go:build dev

package web

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/charmbracelet/log"
)

func HandleWeb(mux *http.ServeMux) {
	viteDevUrl := "http://localhost:5173"

	log.Info("in dev mode, proxying to vite dev server")
	viteProxy := newReverseProxy(viteDevUrl)
	mux.Handle("/", viteProxy)
}

func newReverseProxy(target string) *httputil.ReverseProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL for proxy: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	// Optional: Modify request/response headers if needed
	// proxy.Director = func(req *http.Request) { ... }
	return proxy
}
