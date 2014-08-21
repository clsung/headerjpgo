package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	dstFile = flag.String("to", "a.jpg", "the destination")
	srcFile = flag.String("from", "logo.jpg", "the source jpeg")
	file    = flag.String("file", "hello.txt", "the source text")
)

func main() {
	flag.Parse()
	if !strings.HasSuffix(*srcFile, ".jpg") {
		log.Printf("skipping non .jpg file %v", *srcFile)
		return
	}
	log.Printf("adding %s with file %s and *s", *dstFile, *file, *srcFile)
	bytesJpg, err := ioutil.ReadFile(*srcFile)
	if err != nil {
		log.Fatal(err)
	}
	bytesTxt, err := ioutil.ReadFile(*file)
	length := len(bytesTxt)
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, int16(length))
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(*dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write([]byte{255, 216, 255, 254})
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(bytesTxt)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(bytesJpg[2:])
	if err != nil {
		log.Fatal(err)
	}
}
