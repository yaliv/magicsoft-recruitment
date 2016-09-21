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

type Filename string
type Parent string

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
			// Compare file content.
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

// Seperate path into parent path and filename.
func seperatePath(path string) (Parent, Filename) {
	// Get the index of the last separator.
	sepIdx := strings.LastIndex(path, sep)

	// Get the parent path and the filename.
	var parent Parent
	if sepIdx == 0 {
		parent = Parent("root")
	} else {
		parent = Parent(path[:sepIdx])
	}
	filename := Filename(path[sepIdx+1:])

	return parent, filename
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
