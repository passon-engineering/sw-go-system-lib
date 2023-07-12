package file

import (
	"io/ioutil"
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
