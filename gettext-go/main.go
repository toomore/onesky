package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	//"code.google.com/p/gettext-go/gettext/po"
	"github.com/toomore/gettext-go/gettext/po"
)

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
	csvtopo("onesky.csv")
	//csvdata, _ := readCSV("onesky.csv")
	//createPO("onesky.po", csvdata, 0)
}
