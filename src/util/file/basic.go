package file

import (
	"os"
	"strings"
)

func DeleteFile(path string) error {
	return os.Remove(path)
}

func GetFileNameFromPath(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx == -1 {
		return path
	}
	if idx == len(path)-1 {
		return ""
	}
	return path[idx+1:]
}

func WriteFile(path string, content string) error {
	// make sure the directory exists
	last_idx := strings.LastIndex(path, "/")
	if last_idx > 0 {
		dir := path[:last_idx]
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
