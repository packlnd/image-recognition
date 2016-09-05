package main

import (
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
)

func main() {
	// http://stats.stackexchange.com/questions/39037/how-does-neural-network-recognise-images
}

func train() {
	files, _ := ioutil.ReadDir("./images/good")
	for _, f := range files {
		infile, _ := os.Open(f)
		src, _, _ := image.Decode(infile)
		fmt.Println(src)
	}
}
