package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	//"code.google.com/p/gettext-go/gettext/po"
	"github.com/toomore/gettext-go/gettext/po"
)

func createPO(filename string, csvdata [][]string, rownum int, basedir string) {
	var lang = csvdata[0][rownum]
	var filepath = fmt.Sprintf("%s/%s/%s", basedir, lang, filename)

	os.MkdirAll(fmt.Sprintf("%s/%s", basedir, lang), 0776)

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

func csvtopo(filename, outputdir string) {
	filename = filepath.Base(filename)
	csvdata, _ := readCSV(filename)
	orgfilename := strings.Split(filename, ".")

	for i, _ := range csvdata[0] {
		createPO(fmt.Sprintf("%s.po", orgfilename[0]), csvdata, i, outputdir)
	}
}

var csvpath = flag.String("csv", "", "The paht of csv file.")
var outputdir = flag.String("out", "", "The paht of output po dir.")

func main() {
	flag.Parse()
	if *csvpath == "" {
		flag.PrintDefaults()
		return
	}

	if *outputdir == "" {
		*outputdir = strconv.FormatInt(time.Now().Unix(), 10)
	}
	csvtopo(*csvpath, *outputdir)
	//csvdata, _ := readCSV("onesky.csv")
	//createPO("onesky.po", csvdata, 0)
}
