package main

import (
	"fmt"
	"os"
)

var (
	FILENAME = "examples/stego/example"
	FILETYPE = "pdf"
	SIGNTYPE = "base64"
	OUTPUT   = "embedded_file"
	SIG_MP4  = []byte{
		'\x00', '\x00', '\x00', '\x18', '\x66', '\x74', '\x79', '\x70',
	}

	SIG_PNG = []byte{
		'\x89', '\x50', '\x4E', '\x47', '\x0D', '\x0A', '\x1A', '\x0A',
	}

	SIG_PDF = []byte{
		'\x25', '\x50', '\x44', '\x46', '\x2d',
	}

	SIG_BASE64 = []byte{
		'\x62', '\x61', '\x73', '\x65', '\x36', '\x34', '\x3a',
	}

	SIGN_SEQUENCE = fmt.Sprintf("%x", SIG_MP4)
)

func main() {
	filename := FILENAME + "." + FILETYPE
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}

	firstFile := 0
	if FILETYPE == SIGNTYPE {
		firstFile = 1
	}

	for key, eachByte := range buf {
		if eachByte == SIGNATURE[0] && firstFile == 1 {
			firstFile++
			continue
		}
		if eachByte == SIGNATURE[0] && firstFile != 1 {
			byteSlice := buf[key : key+len(SIGNATURE)]
			sequence := fmt.Sprintf("%x", byteSlice)
			if string(sequence) == SIGN_SEQUENCE {
				fmt.Println("Done,", OUTPUT+"."+SIGNTYPE)
				outputFile := buf[key:]
				newFile, err := os.Create(OUTPUT + "." + SIGNTYPE)
				if err != nil {
					panic(err)
				}
				newFile.Write(outputFile)
				newFile.Close()
				os.Exit(1)
			}
		}
	}
	fmt.Println("Not found!")
}
