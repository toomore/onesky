package main

import (
	"crypto/md5"
	"fmt"
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

func main() {
	fmt.Println(RenderAuth())
}
