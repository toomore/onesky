package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	//"code.google.com/p/gettext-go/gettext/po"
	"github.com/toomore/gettext-go/gettext/po"
)

var poherder = &po.Header{
	ProjectIdVersion:        "PROJECTIDVERSION",
	ContentType:             "text/plain; charset=UTF-8",
	ContentTransferEncoding: "8bit",
	LanguageTeam:            "LANGUAGE-TEAM",
	MimeVersion:             "1.0",
}

func createPO(filename string, csvdata [][]string, rownum int, basedir string) {
	var popath = filepath.Join(basedir, csvdata[0][rownum], filename)

	os.MkdirAll(filepath.Join(basedir, csvdata[0][rownum]), 0776)

	if _, err := os.Stat(popath); os.IsNotExist(err) {
		f, err := os.Create(popath)
		if err != nil {
			log.Fatal(err)
		}
		f.Chmod(0776)
		f.Close()
	}
	pofile, err := po.Load(popath)
	if err != nil {
		log.Fatal(err)
	}

	poherder.Language = csvdata[0][rownum]
	pofile.MimeHeader = *poherder

	for _, v := range csvdata[1:] {
		pofile.Messages = append(pofile.Messages,
			po.Message{
				MsgId:  v[0],
				MsgStr: v[rownum],
			})
	}
	pofile.Save(popath)
}

func readCSV(filename string) ([][]string, error) {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	return reader.ReadAll()
}

func csvtopo(filename, outputdir string) {
	filename = filepath.Base(filename)
	csvdata, err := readCSV(filename)
	if err != nil {
		log.Fatal(err)
	}

	orgfilename := strings.Split(filename, ".")

	for i := range csvdata[0] {
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
