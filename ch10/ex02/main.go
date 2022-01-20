package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"ex02/archive"       //"github.com/YasufumiNakama/go-study/ch10/ex02/archive"
	_ "ex02/archive/tar" //_ "github.com/YasufumiNakama/go-study/ch10/ex02/archive/tar"
	_ "ex02/archive/zip" //_ "github.com/YasufumiNakama/go-study/ch10/ex02/archive/zip"
)

var (
	file = flag.String("file", "sample.zip", "file")
)

func main() {
	flag.Parse()
	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	defer f.Close()

	_, name, err := archive.ReadArchive(f)
	fmt.Printf("Read %s file\n", name)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
}
