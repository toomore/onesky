package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"path"
	"time"
)

type AuthData struct {
	ApiKey    string
	Timestamp string
	Hashkey   string
}

func renderAuth(now time.Time) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s", now.Unix(), APIKEY))))
}

func RenderAuth() *AuthData {
	now := time.Now()
	return &AuthData{
		ApiKey:    APIKEY,
		Timestamp: fmt.Sprint(now.Unix()),
		Hashkey:   renderAuth(now),
	}
}

type OneskyAPI struct{}

var basepath string = path.Base(APIPATH)

func (o OneskyAPI) GetProjectInfo() {
	urlPath := path.Join(basepath, "projects", PROJECTID, "languages")
	resp, _ := http.Get(urlPath)
	fmt.Println(resp.Body)
	defer resp.Body.Close()
}

func main() {
	data := RenderAuth()
	fmt.Println(data)
	//fmt.Println(url.Values{)
}
