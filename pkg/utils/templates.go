package utils

import "path/filepath"

func GetFilesFromRoute(route string) ([]string, error) {
	files, err := filepath.Glob("./" + route)
	if err != nil {
		return nil, err
	}
	return files, nil
}
