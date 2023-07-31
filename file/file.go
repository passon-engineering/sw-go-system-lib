package file

import (
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
	files, err := os.ReadDir(directory)
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
	files, err := os.ReadDir(directory)
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

// DirectoryStats holds the statistics about a directory,
// including the total count of files, directories, and the total size.
//
// Fields:
//   - FileCount: int - the total number of files
//   - DirectoryCount: int - the total number of directories
//   - TotalSize: int64 - the total size of files
type DirectoryStats struct {
	FileCount      int
	DirectoryCount int
	TotalSize      int64
}

// Returns the total size of files in bytes.
func (ds *DirectoryStats) TotalSizeBytes() float64 {
	return float64(ds.TotalSize)
}

// Returns the total size of files in kilobytes.
func (ds *DirectoryStats) TotalSizeKB() float64 {
	return float64(ds.TotalSize) / 1024
}

// Returns the total size of files in megabytes.
func (ds *DirectoryStats) TotalSizeMB() float64 {
	return float64(ds.TotalSize) / (1024 * 1024)
}

// Returns the total size of files in gigabytes.
func (ds *DirectoryStats) TotalSizeGB() float64 {
	return float64(ds.TotalSize) / (1024 * 1024 * 1024)
}

// Returns the total size of files in terabytes.
func (ds *DirectoryStats) TotalSizeTB() float64 {
	return float64(ds.TotalSize) / (1024 * 1024 * 1024 * 1024)
}

// Iteratively traverses the directory structure rooted at 'root',
// counts the number of files and directories up to the specified 'maxDepth', and calculates
// the total size of all files encountered. Directories and files at depths greater than 'maxDepth'
// are ignored.
//
// Parameters:
//   - root: string - the path to the root directory
//   - maxDepth: int - the maximum depth to traverse
//
// Returns:
//   - *DirectoryStats: if the traversal was successful, a pointer to a DirectoryStats struct
//     containing the counts and total size. If the traversal was not successful, nil is returned.
//   - error: if there was an error reading a directory or getting file information, the function
//     returns this error. Otherwise, it returns nil.
//
// Example usage:
//
//	stats, err := countFilesAndFolders("/path/to/your/directory", 2)
//	if err != nil {
//	  fmt.Println("Error:", err)
//	  return
//	}
//	fmt.Printf("Number of files: %d\n", stats.FileCount)
//	fmt.Printf("Number of directories: %d\n", stats.DirectoryCount)
//	fmt.Printf("Total size: %.2f bytes\n", stats.TotalSizeBytes())
func CountFilesAndFolders(path string, maxDepth int) (DirectoryStats, error) {
	stats := DirectoryStats{}

	var walk func(string, int) error
	walk = func(path string, depth int) error {
		if depth > maxDepth {
			return nil
		}

		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				stats.DirectoryCount++
				err := walk(filepath.Join(path, entry.Name()), depth+1)
				if err != nil {
					return err
				}
			} else {
				stats.FileCount++
			}
		}

		return nil
	}

	err := walk(path, 0)
	if err != nil {
		return DirectoryStats{}, err
	}

	return stats, nil
}
