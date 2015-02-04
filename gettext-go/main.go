package main

import (
	"fmt"
	"log"

	"code.google.com/p/gettext-go/gettext/po"
)

func main() {
	pofile, err := po.Load("test.po")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", pofile)
}
