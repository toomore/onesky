package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

//var basepath string = path.Base(APIPATH)

func (o OneskyAPI) httpGet(urlPath string) {
	resp, _ := http.Get(urlPath)
	if content, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Printf("%s", content)
	}
	defer resp.Body.Close()
}

func (o OneskyAPI) GetProjectInfo(params *AuthData) {
	urlParams := params.ToURLValue()
	urlPath := fmt.Sprintf("%s%s?%s", APIPATH, path.Join("projects", PROJECTID, "languages"), urlParams.Encode())
	o.httpGet(urlPath)
}

func (o OneskyAPI) GetFilesList(params *AuthData) {
	urlParams := params.ToURLValue()
	urlPath := fmt.Sprintf("%s%s?%s", APIPATH, path.Join("projects", PROJECTID, "files"), urlParams.Encode())
	o.httpGet(urlPath)
}

func main() {
	data := RenderAuth()
	fmt.Println(data)
	o := OneskyAPI{}
	o.GetProjectInfo(data)
	o.GetFilesList(data)
}
