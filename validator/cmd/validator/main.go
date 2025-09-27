package main

import (
	validator "digito-platform/validator/pkg"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	file := flag.String("file", "", "CMS/PDF file to validate")
	flag.Parse()
	data, _ := ioutil.ReadFile(*file)
	attrs, err := validator.ParseCMS(data)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}
	fmt.Println("Parsed attributes:", len(attrs))
}
