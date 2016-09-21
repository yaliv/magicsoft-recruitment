// Utility: hash file in MD5.
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func hash_file_md5(path string) (string, error) {
	// Initialize variable returnMD5String now in case an error has to be returned.
	var md5String string

	// Check if the path is a directory.
	pathInfo, err := os.Stat(path)
	if err != nil {
		return md5String, err
	}
	if pathInfo.IsDir() {
		return md5String, fmt.Errorf("Cannot hash a dir.")
	}

	// Open the passed argument and check for any error.
	file, err := os.Open(path)
	if err != nil {
		return md5String, err
	}

	// Tell the program to call the following function when the current function returns.
	defer file.Close()

	// Open a new hash interface to write to.
	hash := md5.New()

	// Copy the file in the hash interface and check for any error.
	if _, err := io.Copy(hash, file); err != nil {
		return md5String, err
	}

	// Get the 16 bytes hash.
	hashInBytes := hash.Sum(nil)[:16]

	// Convert the bytes to a string.
	md5String = hex.EncodeToString(hashInBytes)

	return md5String, nil
}
