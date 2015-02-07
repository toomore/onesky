package main

import (
	"encoding/csv"
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

func createPO(filename string, csvdata [][]string, rownum int) {
	var lang = csvdata[0][rownum]
	var filepath = fmt.Sprintf("%s/%s", lang, filename)

	os.Mkdir(fmt.Sprintf("./%s", lang), 0776)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		f, _ := os.Create(filepath)
		f.Chmod(0776)
		f.Close()
	}
	pofile, _ := po.Load(filepath)
	header := po.Header{
		ProjectIdVersion: "Toomore",
	}
	pofile.MimeHeader = header
	for _, v := range csvdata[1:] {
		pofile.Messages = append(pofile.Messages,
			po.Message{
				MsgId:  v[0],
				MsgStr: v[rownum],
			})
	}

	pofile.Save(filepath)
}

func readCSV(filename string) ([][]string, error) {
	csvfile, _ := os.Open(filename)
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	return reader.ReadAll()
}

func csvtopo(filename string) {
	csvdata, _ := readCSV(filename)
	orgfilename := strings.Split(filename, ".")

	for i, _ := range csvdata[0] {
		createPO(fmt.Sprintf("%s.po", orgfilename[0]), csvdata, i)
	}
}

func main() {
	//createWithHeader("test2.po")
	csvtopo("onesky.csv")
	//csvdata, _ := readCSV("onesky.csv")
	//createPO("onesky.po", csvdata, 0)
}
