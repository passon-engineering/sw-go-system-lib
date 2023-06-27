package web

import "testing"

func TestGetContentType(t *testing.T) {
	tests := []struct {
		path          string
		expectedType  string
	}{
		{path: "styles.css", expectedType: "text/css"},
		{path: "script.js", expectedType: "application/javascript"},
		{path: "index.html", expectedType: "text/html"},
		{path: "image.png", expectedType: "image/png"},
		{path: "picture.jpeg", expectedType: "image/jpeg"},
		{path: "photo.jpg", expectedType: "image/jpeg"},
		{path: "animation.gif", expectedType: "image/gif"},
		{path: "data.json", expectedType: "application/json"},	
		{path: "unknown.xyz", expectedType: ""},
	}

	for _, test := range tests {
		contentType := GetContentType(test.path)
		if contentType != test.expectedType && test.expectedType != "" {
			t.Errorf("For path %s, expected content type %s, but got %s", test.path, test.expectedType, contentType)
		}		
	}
}
