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
		f.Close()
	}
	pofile, _ := po.Load(filepath)
	header := po.Header{
		ProjectIdVersion: "Toomore",
	}
	pofile.MimeHeader = header
	//pofile.Messages = append(pofile.Messages, po.Message{
	//	MsgId: "Toomore", MsgStr: "MsgToomore"})

	pofile.Save(filepath)
}

func readCSV(filename string) ([][]string, error) {
	csvfile, _ := os.Open(filename)
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	return reader.ReadAll()
}

func loopCSV(filename string) {
	csvdata, _ := readCSV(filename)
	//orgfilename := strings.Split(filename, ".")

	// gen pot(base: zh_TW).

	for i, langs := range csvdata[0] {
		fmt.Printf("%d [%s]\n", i, langs)
		for ri, value := range csvdata {
			fmt.Println(ri, value[0], "//", value[i])
		}
	}
}

func main() {
	//createWithHeader("test2.po")
	//loopCSV("onesky.csv")
	csvdata, _ := readCSV("onesky.csv")
	createPO("onesky.po", csvdata, 0)
}
