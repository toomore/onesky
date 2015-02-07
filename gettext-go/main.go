package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	//"code.google.com/p/gettext-go/gettext/po"
	"github.com/toomore/gettext-go/gettext/po"
)

func copyAndAdd() {
	pofile, err := po.Load("test.po")
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range pofile.Messages {
		fmt.Println(i, v.MsgId)
	}
	fmt.Println(pofile.MimeHeader)
	n := &po.Message{
		MsgId:  "Toomore",
		MsgStr: "MsgToomore",
	}
	pofile.Messages = append(pofile.Messages, *n)
	for i, v := range pofile.Messages {
		fmt.Println(i, v.MsgId, v.MsgStr)
	}
	pofile.Save("test_result.po")
}

func createWithHeader(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f, _ := os.Create(filename)
		f.Close()
	}
	pofile, _ := po.Load(filename)
	header := po.Header{
		ProjectIdVersion: "Toomore",
	}
	pofile.MimeHeader = header
	pofile.Messages = append(pofile.Messages, po.Message{
		MsgId: "Toomore", MsgStr: "MsgToomore"})

	fmt.Println(header)
	orgfilename := strings.Split(filename, ".")
	pofile.Save(fmt.Sprintf("%s_result.%s", orgfilename[0], orgfilename[1]))
}

func main() {
	createWithHeader("test2.po")
}
