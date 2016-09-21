// Compare two directories.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	sep            string
	srcDir, trgDir string
	Paths          []string
	srcMap, trgMap map[string]bool
)

func main() {
	// Get the source and target directory.
	flag.Parse()
	srcDir = flag.Arg(0)
	trgDir = flag.Arg(1)

	//fmt.Println("Source:", srcDir)
	//fmt.Println("Target:", trgDir)

	// Check if the source exist.
	src, err := os.Stat(srcDir)
	isFatal("Check source:", err)
	// Check if the source is a directory.
	if !src.IsDir() {
		log.Fatalln("Source is not a directory.")
	}

	// Check if the target exist.
	trg, err := os.Stat(trgDir)
	isFatal("Check target:", err)
	// Check if the target is a directory.
	if !trg.IsDir() {
		log.Fatalln("Target is not a directory.")
	}

	//log.Println("[OK] The source and target dir exist.\n")

	// Get the OS file separator.
	sep = string(os.PathSeparator)

	// Prepare the source and target map.
	srcMap = make(map[string]bool)
	trgMap = make(map[string]bool)

	// Walk the source dir.
	filepath.Walk(srcDir, visitSrc)
	filepath.Walk(trgDir, visitTrg)

	//log.Println("[OK] Source and target dir walked successfully.\n")

	// Sort the paths.
	sort.Strings(Paths)

	// Compare the paths.
	for _, path := range Paths {
		//fmt.Println(path)

		var srcOK, trgOK bool

		if _, ok := srcMap[path]; ok {
			srcOK = true
		}
		if _, ok := trgMap[path]; ok {
			trgOK = true
		}

		// If the path exists in the source dir,
		// but not exist in the target dir.
		if srcOK && !trgOK {
			fmt.Println(path, "NEW")
		}

		// If the path exists in the target dir,
		// but not exist in the source dir.
		if trgOK && !srcOK {
			fmt.Println(path, "DELETED")
		}

		// If the path exists both in the source dir and the target dir.
		if srcOK && trgOK {
			// Get their full path.
			srcFullPath := srcDir + path
			trgFullPath := trgDir + path
			//fmt.Printf("Source: %v\nTarget: %v\n_\n", srcFullPath, trgFullPath)

			// Hash the source path.
			srcMd5, err := hash_file_md5(srcFullPath)
			if err != nil {
				if err.Error() != "Cannot hash a dir." {
					log.Printf("Error when hashing %v: %v\n", srcFullPath, err)
				}
				continue
			}

			// Hash the target path.
			trgMd5, err := hash_file_md5(trgFullPath)
			if err != nil {
				if err.Error() != "Cannot hash a dir." {
					log.Printf("Error when hashing %v: %v\n", trgFullPath, err)
				}
				continue
			}

			//fmt.Printf("Source MD5: %v\nTarget MD5: %v\n_\n", srcMd5, trgMd5)

			// Compare both hashes.
			if srcMd5 != trgMd5 {
				fmt.Println(path, "MODIFIED")
			}
		}
	}
}

// Check if fatal error exist.
func isFatal(message string, err error) {
	if err != nil {
		log.Fatalln(message, err)
	}
}

// Append path to the slice, while maintain uniqueness.
func appendPath(path string) {
	for _, v := range Paths {
		if path == v {
			return
		}
	}
	Paths = append(Paths, path)
}

// It is called in every source path found.
func visitSrc(path string, f os.FileInfo, err error) error {
	isFatal("Read a source path:", err)

	if path != srcDir {
		// Trim the root dir from the path.
		path = strings.TrimPrefix(path, srcDir)
		//fmt.Println("Source:", path)

		// Append path to the slice.
		appendPath(path)

		// Add the path to the map.
		srcMap[path] = true
	}

	return nil
}

// It is called in every target path found.
func visitTrg(path string, f os.FileInfo, err error) error {
	isFatal("Read a target path:", err)

	if path != trgDir {
		// Trim the root dir from the path.
		path = strings.TrimPrefix(path, trgDir)
		//fmt.Println("Target:", path)

		// Append path to the slice.
		appendPath(path)

		// Add the path to the map.
		trgMap[path] = true
	}

	return nil
}
