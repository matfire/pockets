//go:build !dev

package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

//go:embed all:dist
var embeddedFiles embed.FS

// registerFrontendHandlers provides the production implementation.
func HandleWeb(mux *http.ServeMux) {
	log.Println("Running without DEV tag. Serving embedded frontend.")

	subFS, err := fs.Sub(embeddedFiles, "frontend/dist")
	if err != nil {
		log.Fatalf("Failed to create sub VFS for embedded files: %v", err)
	}
	staticFS := http.FS(subFS)
	fileServer := http.FileServer(staticFS)

	// Handle all other requests by serving static files or index.html
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := strings.TrimPrefix(path.Clean(r.URL.Path), "/")

		// Check if the requested file exists in the embedded FS
		f, err := staticFS.Open(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				// File doesn't exist, serve index.html for SPA routing
				log.Printf("File %s not found, serving index.html", filePath)
				http.ServeFileFS(w, r, staticFS, "index.html")
				return
			}
			// Other error opening file
			log.Printf("Error opening file %s: %v", filePath, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		f.Close() // Close the file handle after checking existence

		// File exists, let the FileServer handle it
		log.Printf("Serving existing file: %s", filePath)
		fileServer.ServeHTTP(w, r)
	})
}
