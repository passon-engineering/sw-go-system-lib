package web

import (
	"net/http"
	"os"
	"strconv"
)

// Serves a static file as the HTTP response.
//
// It reads the contents of the file specified by the given file path,
// sets the appropriate headers, and writes the file contents to the response.
// If the file does not exist, it returns a "404 Not Found" response.
//
// Parameters:
//  - w: http.ResponseWriter - the response writer to write the HTTP response
//  - r: *http.Request - the HTTP request object
//  - filePath: string - the path to the static file to be served
//
// Returns:
//  - error: an error if any occurred during file reading or writing, or nil if successful
//
// Headers set:
//  - Content-Type: based on the file extension, determined by the GetContentType function
//  - Access-Control-Allow-Origin: allows requests from any origin
//  - Content-Length: the length of the file contents in bytes
//
// Example usage:
//  serveStaticFile(w, r, "/path/to/file.txt")
func ServeStaticFile(w http.ResponseWriter, r *http.Request, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return nil
		}
		return err
	}

	// Set the content type based on the file extension
	contentType := GetContentType(r.URL.Path)
	w.Header().Set("Content-Type", contentType)
		
	// Set additional headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))

	// Write the file contents to the response
	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}
