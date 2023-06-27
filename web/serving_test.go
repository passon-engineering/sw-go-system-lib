package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeStaticFile(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "serving_test.json", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	recorder := httptest.NewRecorder()

	// Call the serveStaticFile function
	err = ServeStaticFile(recorder, req, "serving_test.json")
	if err != nil {
		t.Errorf("serveStaticFile returned an error: %v", err)
	}

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Perform additional assertions on the response if needed
	// ...

	// Example: Check the Content-Type header
	expectedContentType := "application/json"
	actualContentType := recorder.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type header %q, got %q", expectedContentType, actualContentType)
	}

	// Example: Check the Content-Length header
	expectedContentLength := "15" // Example value, replace with the actual expected length
	actualContentLength := recorder.Header().Get("Content-Length")
	if actualContentLength != expectedContentLength {
		t.Errorf("Expected Content-Length header %q, got %q", expectedContentLength, actualContentLength)
	}

	// Example: Check the response body
	expectedResponseBody := "{\"test\":\"test\"}" // Example value, replace with the actual expected body
	actualResponseBody := recorder.Body.String()
	if actualResponseBody != expectedResponseBody {
		t.Errorf("Expected response body %q, got %q", expectedResponseBody, actualResponseBody)
	}

	// Example: Test scenario for a non-existent file
	// Create a mock HTTP request
	notExistReq, err := http.NewRequest("GET", "not_existent.json", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	notExistRecorder := httptest.NewRecorder()

	// Call the serveStaticFile function with a non-existent file
	err = ServeStaticFile(notExistRecorder, notExistReq, "not_existent.json")
	if err != nil {
		t.Errorf("serveStaticFile returned an error: %v", err)
	}

	// Check the response status code (should be 404)
	if notExistRecorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, notExistRecorder.Code)
	}

}
