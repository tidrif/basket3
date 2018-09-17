package main

import (
	"crypto/sha256"
	"fmt"
	"html"
	"mime"
)

func main() {
	var helloWorld string = "<h1>hellooö, world</h1>";
	fmt.Printf("%s%s%s", "Original: ",helloWorld, "\n")
	var helloWorldUTF string = mime.QEncoding.Encode("utf-8", html.UnescapeString(helloWorld))
	fmt.Printf("%s%s%s", "UTF: ",helloWorldUTF, "\n")
	h := sha256.New()
	h.Write([]byte(helloWorldUTF))
	fmt.Printf("%s%x%s", "Encoded: ", h.Sum(nil), "\n")
}