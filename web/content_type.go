package web

import (
	"net/http"
	"path/filepath"
	"strings"
)

var mimeTypes = map[string]string{
	".css":  "text/css",
	".js":   "application/javascript",
	".html": "text/html",
	".png":  "image/png",
	".jpeg": "image/jpeg",
	".jpg":  "image/jpeg",
	".gif":  "image/gif",
	".json": "application/json",
}

// Returns the MIME type based on the file extension.
//
// It retrieves the file extension from the given path, converts it to lowercase,
// and checks if it exists in the mimeTypes map. If a matching MIME type is found,
// it is returned. Otherwise, http.DetectContentType is used to determine the
// content type dynamically based on the file's content.
//
// Parameters:
//   - path: string - the path to the file
//
// Returns:
//   - string: the MIME type of the file
func GetContentType(path string) string {
	extension := strings.ToLower(filepath.Ext(path))
	if mimeType, ok := mimeTypes[extension]; ok {
		return mimeType
	}
	return http.DetectContentType(nil)
}
