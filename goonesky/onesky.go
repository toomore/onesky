package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"reflect"
	"time"
)

type AuthData struct {
	ApiKey    string `json:"api_key"`
	Timestamp string `json:"timestamp"`
	Hashkey   string `json:"dev_hash"`
}

func renderAuth(timestamp string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s", timestamp, APISECRET))))
}

func RenderAuth() *AuthData {
	timestamp := fmt.Sprint(time.Now().Unix())
	return &AuthData{
		ApiKey:    APIKEY,
		Timestamp: timestamp,
		Hashkey:   renderAuth(timestamp),
	}
}

func (auth AuthData) ToURLValue() url.Values {
	urlParams := url.Values{}
	p := reflect.ValueOf(auth)
	for i := 0; i < p.NumField(); i++ {
		f := p.Field(i)
		tagName := p.Type().Field(i).Tag.Get("json")
		urlParams.Add(tagName, fmt.Sprint(f))
	}
	return urlParams
}

type OneskyAPI struct{}

var basepath, _ = url.Parse(APIPATH)

func (o OneskyAPI) httpGet(urlPath string) {
	resp, _ := http.Get(urlPath)
	if content, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Printf("%s", content)
	}
	defer resp.Body.Close()
}

func (o OneskyAPI) httpPostForm(urlPath string, data url.Values) {
	resp, _ := http.PostForm(urlPath, data)
	if content, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Printf("%s", content)
	}
	defer resp.Body.Close()
}

func (o OneskyAPI) httpPostData(urlPath string, data *os.File) {
	//resp, _ := http.Post(urlPath, "multipart/form-data", strings.NewReader(data.Encode()))
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", data.Name())
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.Copy(fw, data); err != nil {
		return
	}
	w.Close()

	resp, _ := http.Post(urlPath, w.FormDataContentType(), bytes.NewReader(b.Bytes()))
	if content, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Printf("%s", content)
	}
	defer resp.Body.Close()
}

func (o OneskyAPI) UploadPO(params *AuthData, files ...string) {
	urlParams := params.ToURLValue()
	urlParams.Add("file_format", "GNU_PO")
	basepath.Path = path.Join(APIVERSION, "projects", PROJECTID, "files")
	basepath.RawQuery = urlParams.Encode()

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("`%s` not find.", filename)
		}
		o.httpPostData(basepath.String(), file)
	}
}

func (o OneskyAPI) GetProjectInfo(params *AuthData) {
	basepath.Path = path.Join(APIVERSION, "projects", PROJECTID, "languages")
	basepath.RawQuery = params.ToURLValue().Encode()
	o.httpGet(basepath.String())
}

func (o OneskyAPI) GetFilesList(params *AuthData) {
	basepath.Path = path.Join(APIVERSION, "projects", PROJECTID, "files")
	basepath.RawQuery = params.ToURLValue().Encode()
	o.httpGet(basepath.String())
}

func main() {
	auth := RenderAuth()
	fmt.Println(auth)
	o := OneskyAPI{}
	o.GetProjectInfo(auth)
	o.GetFilesList(auth)

	o.UploadPO(auth, "onesky.po", "test.po")
}
