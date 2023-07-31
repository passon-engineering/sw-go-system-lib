package file

import (
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteAllExceptIgnored(t *testing.T) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	ioutil.WriteFile(filepath.Join(dir, "file1"), []byte("file1"), 0777)
	ioutil.WriteFile(filepath.Join(dir, "file2"), []byte("file2"), 0777)

	err = DeleteAllExceptIgnored(dir, map[string]bool{"file1": true})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filepath.Join(dir, "file1")); os.IsNotExist(err) {
		t.Fatal("File1 should not be deleted")
	}

	if _, err := os.Stat(filepath.Join(dir, "file2")); !os.IsNotExist(err) {
		t.Fatal("File2 should be deleted")
	}
}

func TestDeleteAll(t *testing.T) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	ioutil.WriteFile(filepath.Join(dir, "file1"), []byte("file1"), 0777)
	ioutil.WriteFile(filepath.Join(dir, "file2"), []byte("file2"), 0777)

	err = DeleteAll(dir)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the directory is empty
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 0 {
		t.Fatal("Directory should be empty")
	}
}

func TestDelete(t *testing.T) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	file := filepath.Join(dir, "file1")
	ioutil.WriteFile(file, []byte("file1"), 0777)

	err = Delete(file)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(file); !os.IsNotExist(err) {
		t.Fatal("File should be deleted")
	}
}

func TestCountFilesAndFolders(t *testing.T) {
	// Setup a test directory with some files and directories
	tmpDir, err := ioutil.TempDir("", "test-dir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %s", err)
	}
	defer os.RemoveAll(tmpDir) // clean up

	subDir := filepath.Join(tmpDir, "subdir")
	err = os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdir: %s", err)
	}

	for _, f := range []string{"file1", "file2", filepath.Join("subdir", "file3")} {
		err = ioutil.WriteFile(filepath.Join(tmpDir, f), []byte("test"), 0644)
		if err != nil {
			t.Fatalf("Failed to write file: %s", err)
		}
	}

	// Test the function with maxDepth 0
	stats, err := CountFilesAndFolders(tmpDir, 0)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if stats.FileCount != 2 || stats.DirectoryCount != 1 {
		t.Errorf("Expected FileCount and DirectoryCount to be 2 and 1, got %d and %d", stats.FileCount, stats.DirectoryCount)
	}

	// Test the function with maxDepth 1
	stats, err = CountFilesAndFolders(tmpDir, 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if stats.FileCount != 3 || stats.DirectoryCount != 1 {
		t.Errorf("Expected FileCount and DirectoryCount to be 3 and 1, got %d and %d", stats.FileCount, stats.DirectoryCount)
	}

	// Test the function with maxDepth 2
	stats, err = CountFilesAndFolders(tmpDir, 2)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if stats.FileCount != 3 || stats.DirectoryCount != 1 {
		t.Errorf("Expected FileCount and DirectoryCount to be 3 and 1, got %d and %d", stats.FileCount, stats.DirectoryCount)
	}

	// Test the function with maxDepth math.MaxInt64
	stats, err = CountFilesAndFolders(tmpDir, math.MaxInt64)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if stats.FileCount != 3 || stats.DirectoryCount != 1 {
		t.Errorf("Expected FileCount and DirectoryCount to be 3 and 1, got %d and %d", stats.FileCount, stats.DirectoryCount)
	}

}
