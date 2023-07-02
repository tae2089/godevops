package file

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func deleteOldFiles(dirPath string, daysOld int) error {
	cutoffTime := time.Now().AddDate(0, 0, -daysOld)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() && info.ModTime().Before(cutoffTime) {
			fmt.Printf("Deleting file: %s\n", path)
			if err := os.Remove(path); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
