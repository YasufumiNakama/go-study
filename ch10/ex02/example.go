package main

import (
	"archive/tar"
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	file = flag.String("file", "sample.zip", "file")
)

func main() {
	flag.Parse()
	file := *file
	format := file[strings.LastIndex(file, "."):]
	switch format {
	case ".zip":
		readZip(file)
	case ".tar":
		readTar(file)
	default:
		fmt.Fprintf(os.Stderr, "file should be .zip or .tar\n")
		os.Exit(1)
	}
}

func readTar(file string) {
	// https://github.com/golang/go/blob/master/src/archive/tar/example_test.go
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	defer f.Close()

	tr := tar.NewReader(f)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}

func readZip(file string) {
	// https://github.com/golang/go/blob/master/src/archive/zip/example_test.go#L50-L76
	// Open a zip archive for reading.
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}
