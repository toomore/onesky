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