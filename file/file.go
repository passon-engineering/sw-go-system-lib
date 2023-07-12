package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Deletes all directories within the given directory, except those whose names
// are specified in the 'ignore' map.
//
// The function reads the contents of the specified directory and iterates over each file
// or subdirectory. If the name of a file or subdirectory is a key in the 'ignore' map,
// the function skips it. Otherwise, if it's a directory, the function removes it. Files
// that are not directories are also ignored.
//
// This function is especially useful for cleaning up a directory while preserving certain
// files or subdirectories.
//
// Parameters:
//   - directory: string - the path of the directory to clean up
//   - ignore: map[string]bool - a map where the keys are the names of files or subdirectories to ignore
//
// Returns:
//   - error: if there was an error reading the directory or deleting a subdirectory, the function
//     returns this error. Otherwise, it returns nil.
//
// Example usage:
//
//	ignoreFiles := map[string]bool{".gitignore": true}
//	err := DeleteAllExceptIgnored("repositories/", ignoreFiles)
//	if err != nil {
//	  panic(err)
//	}
func DeleteAllExceptIgnored(directory string, ignore map[string]bool) error {
	// Read the directory
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, f := range files {
		// Skip ignored files
		if _, ok := ignore[f.Name()]; ok {
			continue
		}

		// Create the full path of the file or directory
		fullPath := filepath.Join(directory, f.Name())

		// If it's a directory or a file, remove it
		err = os.RemoveAll(fullPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// Deletes all directories and files within the given directory.
//
// The function reads the contents of the specified directory and iterates over each file
// or subdirectory. It then removes each one, whether it's a file or a directory.
//
// This function is especially useful for cleaning up a directory completely.
//
// Parameters:
//   - directory: string - the path of the directory to clean up
//
// Returns:
//   - error: if there was an error reading the directory or deleting a subdirectory or a file, the function
//     returns this error. Otherwise, it returns nil.
//
// Example usage:
//
//	err := DeleteAll("repositories/")
//	if err != nil {
//	  panic(err)
//	}
func DeleteAll(directory string) error {
	// Read the directory
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, f := range files {
		// Create the full path of the file or directory
		fullPath := filepath.Join(directory, f.Name())

		// If it's a directory or a file, remove it
		err = os.RemoveAll(fullPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// Deletes the given file or directory and all its contents if it is a directory.
//
// The function removes the specified path, whether it's a file or a directory. If it's a directory,
// all its contents will also be deleted.
//
// This function is useful for removing a specific file or directory.
//
// Parameters:
//   - path: string - the path of the file or directory to remove
//
// Returns:
//   - error: if there was an error deleting the file or directory, the function
//     returns this error. Otherwise, it returns nil.
//
// Example usage:
//
//	err := Delete("repositories/some_file_or_directory")
//	if err != nil {
//	  panic(err)
//	}
func Delete(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
