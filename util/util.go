// Package util is the util(util method) package for all of the openengine
package util

import "os"

// FileExists check if the input file exist.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
