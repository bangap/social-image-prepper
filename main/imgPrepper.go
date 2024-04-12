package main

import (
	_ "image/png"
)

var f = "../testdata/test.png"

func main() {
	// open file to r

	decodedImg := open(f)

	convertToPng(decodedImg)

}
